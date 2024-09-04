var log = logrus.New()

func Info(msg string, fields logrus.Fields) {
    log.WithFields(fields).Info(msg)
}

// pkg/metrics/prometheus.go
var (
    TasksProcessed = promauto.NewCounter(prometheus.CounterOpts{
        Name: "tasks_processed_total",
        Help: "The total number of processed tasks",
    })
    TaskDuration = promauto.NewHistogram(prometheus.HistogramOpts{
        Name:    "task_duration_seconds",
        Help:    "The duration of task execution in seconds",
        Buckets: prometheus.LinearBuckets(0, 10, 20),
    })
)