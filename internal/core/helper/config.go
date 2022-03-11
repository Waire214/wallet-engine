package helper

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	AppName            string `mapstructure:"app_name"`
	ServiceAddress     string `mapstructure:"service_address"`
	ServicePort        string `mapstructure:"service_port"`
	ServiceMode        string `mapstructure:"service_mode"`
	MONGODB_URI        string `mapstructure:"MONGODB_URI"`
	DbType             string `mapstructure:"db_type"`
	MongoDbHost        string `mapstructure:"mongo_db_host"`
	MongoDbName        string `mapstructure:"mongo_db_name"`
	MongoDbUserName    string `mapstructure:"mongo_db_username"`
	MongoDbPassword    string `mapstructure:"mongo_db_password"`
	MongoDbPort        string `mapstructure:"mongo_db_port"`
	MongoDbAuthDb      string `mapstructure:"mongo_db_auth_db"`
	Server             string `mapstructure:"server"`
	ServiceName        string `mapstructure:"service_name"`
	LogFile            string `mapstructure:"log_file"`
	LogDir             string `mapstructure:"log_dir"`
	ExternalConfigPath string `mapstructure:"external_config_path"`
	PageLimit          string `mapstructure:"page_limit"`
}

func loadEnv(path string) (config ConfigStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("wallet")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return ConfigStruct{}, err
	}
	err = viper.Unmarshal(&config)
	return
}

func returnConfig() ConfigStruct {
	config, err := loadEnv(".")
	if err != nil {
		log.Println(err)
	}
	if config.ExternalConfigPath != "" {
		viper.Reset()
		config, err = loadEnv(config.ExternalConfigPath)
		if err != nil {
			log.Println(err)
		}
	}
	return config
}

var Config = returnConfig()
