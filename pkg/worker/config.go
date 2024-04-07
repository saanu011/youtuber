package worker

type Config struct {
	Redis            RedisConfig
	Concurrency      int
	JobRetryMaxCount int `mapstructure:"JOB_RETRY_MAX_COUNT"`
}

type RedisConfig struct {
	Addr string `mapstructure:"ADDRESS"`
}
