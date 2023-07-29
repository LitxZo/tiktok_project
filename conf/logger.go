package conf

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zap.DebugLevel
	if !viper.GetBool("Log.developMode") {
		logMode = zap.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), logMode)
	return zap.New(core).Sugar()

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	rootDir, _ := os.Getwd()
	stSeparator := string(filepath.Separator)
	logPath := rootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"

	lumberjackSync := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    viper.GetInt("maxSize"), // megabytes
		MaxBackups: viper.GetInt("maxBackups"),
		MaxAge:     viper.GetInt("maxAge"), //days
		Compress:   true,                   // disabled by default
	}

	return zapcore.AddSync(lumberjackSync)
}
