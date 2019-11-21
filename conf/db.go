package conf

import (
	"github.com/spf13/viper"
)

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

func GetDBConf() DbConf {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.ReadInConfig()
	var MasterDbConfig DbConf = DbConf{
		Host:   viper.GetString("mysql.host"),
		Port:   viper.GetString("mysql.port"),
		User:   viper.GetString("mysql.user"),
		Pwd:    viper.GetString("mysql.password"),
		DbName: viper.GetString("mysql.db_name"),
	}
	return MasterDbConfig
}
