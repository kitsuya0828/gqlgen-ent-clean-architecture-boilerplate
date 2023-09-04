package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Addr                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
			Charset   string
			Loc       string
		}
	}
	Server struct {
		Address string
	}
}

// C is config variable
var C config

// ReadConfigOption is a config option
type ReadConfigOption struct {
	AppEnv string
}

// ReadConfig configures config file
func ReadConfig(option ReadConfigOption) {
	Config := &C

	if IsDev() {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.dev")
	} else if IsProd() {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.prod")
	} else if IsTest() || (option.AppEnv == Test) {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.test")
	} else if IsE2E() || (option.AppEnv == E2E) {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.e2e")
	} else {
		log.Fatalf("failed to get APP_ENV: %s", os.Getenv("APP_ENV"))
	}

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln(err)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
