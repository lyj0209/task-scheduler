package kafka

import (
    "encoding/json"
    "github.com/Shopify/sarama"
    "github.com/lyj0209/task-scheduler/internal/models"
)

type KafkaQueue struct {
    producer sarama.SyncProducer
    consumer sarama.Consumer
    topic    string
}

func NewKafkaQueue(brokers []string, topic string) (*KafkaQueue, error) {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return nil, err
    }

    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        return nil, err
    }

    return &KafkaQueue{
        producer: producer,
        consumer: consumer,
        topic:    topic,
    }, nil
}

func (k *KafkaQueue) PublishTask(task *models.Task) error {
    taskJSON, err := json.Marshal(task)
    if err != nil {
        return err
    }

    _, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
        Topic: k.topic,
        Value: sarama.StringEncoder(taskJSON),
    })
    return err
}

func (k *KafkaQueue) ConsumeTask() (*models.Task, error) {
    partitionConsumer, err := k.consumer.ConsumePartition(k.topic, 0, sarama.OffsetNewest)
    if err != nil {
        return nil, err
    }
    defer partitionConsumer.Close()

    msg := <-partitionConsumer.Messages()
    var task models.Task
    err = json.Unmarshal(msg.Value, &task)
    if err != nil {
        return nil, err
    }

    return &task, nil
}

func (k *KafkaQueue) Close() error {
    if err := k.producer.Close(); err != nil {
        return err
    }
    return k.consumer.Close()
}