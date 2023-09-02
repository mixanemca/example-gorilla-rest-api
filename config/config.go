package config

import (
	"github.com/gotify/configor"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	App struct {
		ListenAddr string `default:""`
		Port       string `default:"8080"`
	}
	Database struct {
		Host     string `default:"localhost" env:"DB_HOST"`
		Port     string `default:"5432" env:"DB_PORT"`
		Name     string `default:"gapi" env:"DB_NAME"`
		User     string `default:"postgres" env:"DB_USER"`
		Password string `default:"postgres" env:"DB_PASSWORD"`
	}
}

var Cfg *Config

// load from config file(s), just for example
func configFiles() []string {
	return []string{"config.yml"}
}

func InitConfig() {
	InitLogger()

	if Cfg != nil {
		return
	}
	conf := new(Config)
	err := configor.New(&configor.Config{}).Load(conf, configFiles()...)
	if err != nil {
		panic(err)
	}
	Cfg = conf
}

func InitLogger() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:        "2006/01/02 - 15:04:05",
		DisableLevelTruncation: true,
		QuoteEmptyFields:       true,
		FullTimestamp:          true,
	})
}
