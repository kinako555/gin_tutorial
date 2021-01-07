package setting

import (
	"time"
	"log"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	DBHost   string `envconfig:"DBHost"`
	Username string `envconfig:"Username"`
	Password string `envconfig:"Password"`
	Database string `envconfig:"Database"`
}

type ServerConfig struct {
	HTTPPort     int           `envconfig:"PORT"`
	ReadTimeout  time.Duration `envconfig:"READTIMEOUT"`
	WriteTimeout time.Duration `envconfig:"WRITETIMEOUT"`
}

type AppConfig struct {
	RunMode   string `envconfig:"RUNMODE"`
	PageSize  int    `envconfig:"PAGESIZE"`
	JWTSecret string `envconfig:"JWTSECRET"`
}

type config struct {
	DBConfig
	ServerConfig
	AppConfig
}
var Config = config{}

func init() {
	err := envconfig.Process("", &Config)
	if err != nil {
		log.Fatalf("Fail to load config wiht env :%v", err)
	}
}