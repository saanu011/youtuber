package redis

type Config struct {
	Address      string `mapstructure:"ADDRESS"`
	DialTimeout  int    `mapstructure:"DIAL_TIMEOUT"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
	PoolSize     int    `mapstructure:"POOL_SIZE"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
}
