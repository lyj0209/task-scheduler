package mysql

type MySQLStorage struct {
    db *sql.DB
}

func (m *MySQLStorage) CreateTask(task *models.Task) error {
    // 实现将任务保存到MySQL的逻辑
}

// internal/storage/redis/redis.go
type RedisStorage struct {
    client *redis.Client
}

func (r *RedisStorage) UpdateTaskStatus(taskID string, status string) error {
    // 实现更新任务状态到Redis的逻辑
}