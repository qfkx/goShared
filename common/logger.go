package common

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger Zap
// var sugarLogger *zap.SugaredLogger
//  sugarLogger = InitLogger("./test1.log")
//  defer sugarLogger.Sync()
// sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
func InitLogger(filename string) *zap.SugaredLogger {
	// 初始化Logger
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)

	// 获取配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}
