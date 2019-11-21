package conf

import "github.com/spf13/viper"

type RedisConf struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func GetRedisConf() RedisConf {
	return struct {
		Host     string
		Port     int
		Password string
		DB       int
	}{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetInt("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db")}
}
