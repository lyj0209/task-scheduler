package mysql

import (
    "database/sql"
    "encoding/json"
    "github.com/lyj0209/task-scheduler/internal/models"
    _ "github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
    db *sql.DB
}

func NewMySQLStorage(dsn string) (*MySQLStorage, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    return &MySQLStorage{db: db}, nil
}

func (s *MySQLStorage) CreateTask(task *models.Task) error {
    _, err := s.db.Exec("INSERT INTO tasks (type, status) VALUES (?, ?)", task.Type, task.Status)
    return err
}

func (s *MySQLStorage) UpdateTask(task *models.Task) error {
    resultJSON, err := json.Marshal(task.Result)
    if err != nil {
        return err
    }
    _, err = s.db.Exec("UPDATE tasks SET status = ?, result = ? WHERE id = ?", task.Status, resultJSON, task.ID)
    return err
}

func (s *MySQLStorage) GetPendingTasks() ([]*models.Task, error) {
    rows, err := s.db.Query("SELECT id, type FROM tasks WHERE status = 'pending' LIMIT 10")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []*models.Task
    for rows.Next() {
        task := &models.Task{Status: models.TaskStatusPending}
        err := rows.Scan(&task.ID, &task.Type)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}

func (s *MySQLStorage) GetOrderCount24h() (int, error) {
    var count int
    err := s.db.QueryRow("SELECT COUNT(*) FROM orders WHERE created_at > DATE_SUB(NOW(), INTERVAL 24 HOUR)").Scan(&count)
    if err != nil {
        return 0, err
    }
    return count, nil
}

func (s *MySQLStorage) GetHotProducts(limit int) (map[string]int, error) {
    rows, err := s.db.Query(`
        SELECT product_id, COUNT(*) as order_count 
        FROM order_items oi 
        JOIN orders o ON oi.order_id = o.id 
        WHERE o.created_at > DATE_SUB(NOW(), INTERVAL 24 HOUR) 
        GROUP BY product_id 
        ORDER BY order_count DESC 
        LIMIT ?
    `, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    hotProducts := make(map[string]int)
    for rows.Next() {
        var productID string
        var orderCount int
        err := rows.Scan(&productID, &orderCount)
        if err != nil {
            return nil, err
        }
        hotProducts[productID] = orderCount
    }
    return hotProducts, nil
}