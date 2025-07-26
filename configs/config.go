package configs

import (
	"github.com/spf13/viper"
	"os"
)

func Load() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBPassword() string {
	return os.Getenv("DB_PASS")
}
