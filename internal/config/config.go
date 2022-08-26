package config

import (
	"mihoyo-bbs-genshin-sign/internal/util"
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
	Addr                    string `mapstructure:"addr" yaml:"addr"`
	EnablePrometheusMetrics bool   `mapstructure:"enable-prometheus-metrics" yaml:"enable-prometheus-metrics"`
	PrometheusMetricsAddr   string `mapstructure:"prometheus-metrics-addr" yaml:"prometheus-metrics-addr"`
}

var defaultConfig = Config{
	System: systemConf{
		Addr:                    "0.0.0.0:5001",
		EnablePrometheusMetrics: true,
		PrometheusMetricsAddr:   "0.0.0.0:9900",
	},
	Zap: zapConf{
		Level:           "info",
		LogDir:          "sign_log",
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

	var err error
	// init all dir
	if err = util.CreateDirs(DefaultConfigPath, DefaultDataPath); err != nil {
		panic(err)
	}

	// read and load config
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
