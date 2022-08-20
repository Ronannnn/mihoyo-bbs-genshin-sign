package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	System systemConf `mapstructure:"system"`
	Zap    zapConf    `mapstructure:"zap"`
	Db     DbConf     `mapstructure:"db"`
	Jwt    jwtConf    `mapstructure:"jwt"`
}

type systemConf struct {
	Addr string `mapstructure:"addr"`
}

type jwtConf struct {
	Duration int    `mapstructure:"duration"`
	Secret   string `mapstructure:"secret"`
}

func init() {
	// set timezone to utc
	time.Local = time.FixedZone("utc", 0)

	// read and load config
	var err error
	v := viper.New()
	v.SetConfigFile(getConfigFilePath())
	if err = v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err = loadConfig(v); err != nil {
		panic(err)
	}
	// watching and updating Conf without application restart
	v.OnConfigChange(func(e fsnotify.Event) {
		if err = loadConfig(v); err != nil {
			panic(err)
		}
	})
	v.WatchConfig()
}

func loadConfig(v *viper.Viper) (err error) {
	if err = v.Unmarshal(&Conf); err != nil {
		return
	}
	if err = initLogger(Conf.Zap); err != nil {
		return
	}
	Logger.Debug("Log loaded")
	if err = initDb(Conf.Db); err != nil {
		return
	}
	Logger.Debug("Db loaded")
	return
}

// getConfigFilePath get configuration file path
// priority: command line >> environment variable >> default value
func getConfigFilePath() (config string) {
	// from command line
	flag.StringVar(&config, "c", "", "input config file path")
	flag.Parse()
	if config != "" {
		fmt.Println("Config file passing from command line:", config)
		return
	}
	// from env var
	if env := os.Getenv(EnvConfigKey); env != "" {
		config = env
		fmt.Println("Config file passing from environment variable:", config)
		return
	}
	// from default value
	config = DefaultConfigFilename
	fmt.Println("Default config file:", config)
	return
}
