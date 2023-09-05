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
			TLS       string
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

	if IsProd() {
		viper.BindEnv("Database.User", "DATABASE_USER")
		viper.BindEnv("Database.Password", "DATABASE_PASSWORD")
		viper.BindEnv("Database.Net", "DATABASE_NET")
		viper.BindEnv("Database.Addr", "DATABASE_ADDR")
		viper.BindEnv("Database.DBName", "DATABASE_DBNAME")
		viper.BindEnv("Database.AllowNativePasswords", "DATABASE_ALLOWNATIVEPASSWORDS")
		viper.BindEnv("Database.Params.ParseTime", "DATABASE_PARAMS_PARSETIME")
		viper.BindEnv("Database.Params.Charset", "DATABASE_PARAMS_CHARSET")
		viper.BindEnv("Database.Params.Loc", "DATABASE_PARAMS_LOC")
		viper.BindEnv("Database.Params.TLS", "DATABASE_PARAMS_TLS")
		viper.BindEnv("Server.Address", "SERVER_ADDRESS")
	} else {
		if IsDev() {
			viper.AddConfigPath(filepath.Join(rootDir(), "config"))
			viper.SetConfigName("config.dev")
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
