type Config struct {
    MySQL struct {
        Host     string
        Port     int
        Username string
        Password string
        Database string
    }
    Redis struct {
        Host string
        Port int
    }
    Kafka struct {
        Brokers []string
        Topic   string
    }
}

func LoadConfig(path string) (*Config, error) {
    // 使用 Viper 加载配置文件的逻辑
}