package config

import (
	"mihoyo-bbs-genshin-sign/util"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

type zapConf struct {
	Level           zapcore.Level `yaml:"level"`
	Prefix          string        `yaml:"prefix"`
	Directory       string        `yaml:"directory"`
	LinkName        string        `yaml:"link-name"`
	EncodeLevel     string        `yaml:"encode-level"`
	LogInConsole    bool          `yaml:"log-in-console"`
	LogInRotatefile bool          `yaml:"log-in-rotatefile"`
}

var zapConfig = &zapConf{
	Level:           zapcore.InfoLevel,
	Prefix:          "[JRS]",
	Directory:       "logs",
	LinkName:        "latest_log",
	EncodeLevel:     "LowercaseLevelEncoder",
	LogInConsole:    true,
	LogInRotatefile: true,
}

func init() {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(newEncoderConfig()),
		newWriteSyncer(),
		zapConfig.Level,
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
	Logger.Info("Logger initialized")
}

// newWriteSyncer get multiple write syncer according to 'LogInConsole' and 'LogInRotatefile' in conf
func newWriteSyncer() zapcore.WriteSyncer {
	var multiWriter []zapcore.WriteSyncer
	if zapConfig.LogInConsole {
		multiWriter = append(multiWriter, zapcore.AddSync(os.Stdout))
	}
	if zapConfig.LogInRotatefile {
		if createErr := util.CreateDirs(zapConfig.Directory); createErr != nil {
			panic(createErr)
		}
		var fileWriter *zaprotatelogs.RotateLogs
		fileWriter, _ = zaprotatelogs.New(
			path.Join(zapConfig.Directory, "%Y-%m-%d.log"),
			zaprotatelogs.WithLinkName(zapConfig.LinkName),
			zaprotatelogs.WithMaxAge(7*24*time.Hour),
			zaprotatelogs.WithRotationTime(24*time.Hour),
		)
		multiWriter = append(multiWriter, zapcore.AddSync(fileWriter))
	}
	return zapcore.NewMultiWriteSyncer(multiWriter...)
}

func newEncoderConfig() (encoderConf zapcore.EncoderConfig) {
	encoderConf = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     customizedTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	switch zapConfig.EncodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		encoderConf.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		encoderConf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		encoderConf.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return encoderConf
}

func customizedTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(zapConfig.Prefix + " " + "2006-01-02 15:04:05.000"))
}
