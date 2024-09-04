package redis

import (
	"crypto/tls" 
    "github.com/go-redis/redis/v8"
    "context"
    "strconv"
    "time"
)

type RedisStorage struct {
    client *redis.Client
}

func NewRedisStorage(addr string) (*RedisStorage, error) {
    client := redis.NewClient(&redis.Options{
        Addr: addr,
        TLSConfig: &tls.Config{
            InsecureSkipVerify: true,
        },
    })

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := client.Ping(ctx).Result()
    if err != nil {
        return nil, err
    }

    return &RedisStorage{client: client}, nil
}

func (s *RedisStorage) SetOrderCount24h(count int) error {
    return s.client.Set(context.Background(), "order_count_24h", count, 1*time.Hour).Err()
}

func (s *RedisStorage) GetOrderCount24h() (int, error) {
    val, err := s.client.Get(context.Background(), "order_count_24h").Result()
    if err != nil {
        return 0, err
    }
    return strconv.Atoi(val)
}

func (s *RedisStorage) UpdateHotProducts(products map[string]int) error {
    ctx := context.Background()
    pipe := s.client.Pipeline()

    pipe.Del(ctx, "hot_products")
    for product, score := range products {
        pipe.ZAdd(ctx, "hot_products", &redis.Z{Score: float64(score), Member: product})
    }

    _, err := pipe.Exec(ctx)
    return err
}

func (s *RedisStorage) GetHotProducts(limit int) ([]string, error) {
    products, err := s.client.ZRevRange(context.Background(), "hot_products", 0, int64(limit-1)).Result()
    if err != nil {
        return nil, err
    }
    return products, nil
}