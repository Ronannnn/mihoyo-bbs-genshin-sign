package config

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"mihoyo-bbs-genshin-sign/util"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

type zapConf struct {
	Level           string `mapstructure:"level"`
	Directory       string `mapstructure:"directory"`
	Filename        string `mapstructure:"filename"`
	LogInConsole    bool   `mapstructure:"log-in-console"`
	LogInRotatefile bool   `mapstructure:"log-in-rotatefile"`
}

func initLogger(zapConfig zapConf) (err error) {
	var level zapcore.Level
	if level, err = zapcore.ParseLevel(zapConfig.Level); err != nil {
		return
	}
	var writeSyncer zapcore.WriteSyncer
	if writeSyncer, err = newWriteSyncer(zapConfig); err != nil {
		return
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "logger",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(LogTimeFormat))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		writeSyncer,
		level,
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
	return
}

// newWriteSyncer get multiple write syncers
// 1. stdout if LogInConsole is enabled
// 2. RotateLogs if LogInRotatefile is enabled
func newWriteSyncer(zapConfig zapConf) (syncer zapcore.WriteSyncer, err error) {
	var multiWriter []zapcore.WriteSyncer
	if zapConfig.LogInConsole {
		multiWriter = append(multiWriter, zapcore.AddSync(os.Stdout))
	}
	if zapConfig.LogInRotatefile {
		// create directory for storing log files
		if err = util.CreateDirs(zapConfig.Directory); err != nil {
			return
		}
		var fileWriter = &lumberjack.Logger{
			Filename:   zapConfig.Filename,
			MaxSize:    1, // rotate when the size gets 1MB
			MaxBackups: 0, // 0 backup: keep all old files
			MaxAge:     0, // 0 days: keep all old files
		}
		multiWriter = append(multiWriter, zapcore.AddSync(fileWriter))
	}
	return zapcore.NewMultiWriteSyncer(multiWriter...), nil
}
