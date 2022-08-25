package config

import (
	"mihoyo-bbs-genshin-sign/util"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	System systemConf `mapstructure:"system" yaml:"system"`
	Zap    zapConf    `mapstructure:"zap" yaml:"zap"`
	Db     DbConf     `mapstructure:"db" yaml:"db"`
}

type systemConf struct {
	Addr string `mapstructure:"addr" yaml:"addr"`
}

var defaultConfig = Config{
	System: systemConf{
		Addr: "0.0.0.0:5001",
	},
	Zap: zapConf{
		Level:           "info",
		Directory:       "logs",
		Filename:        "latest.log",
		LogInConsole:    true,
		LogInRotatefile: true,
	},
	Db: DbConf{
		DbFilename: "sign.db",
	},
}

func init() {
	// set timezone to utc
	time.Local = time.FixedZone("utc", 0)

	// read and load config
	var err error
	viper.AddConfigPath(DefaultConfigPath)
	viper.SetConfigName(DefaultConfigName)
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			// write to local file
			if err = util.WriteYaml(DefaultConfigFilename, defaultConfig); err != nil {
				panic(err)
			}
			// read config file again after creating
			if err = viper.ReadInConfig(); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	if err = loadConfig(); err != nil {
		panic(err)
	}
	// watching and updating Conf without application restart
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err = loadConfig(); err != nil {
			panic(err)
		}
	})
	viper.WatchConfig()
}

func loadConfig() (err error) {
	if err = viper.Unmarshal(&Conf); err != nil {
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
