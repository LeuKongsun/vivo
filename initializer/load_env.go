package initializer

import (
	"time"

	"github.com/spf13/viper"
)


type Config struct {
	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBPort string `mapstructure:"POSTGRES_PORT"`
	DBUser string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_DB"`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}