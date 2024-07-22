package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./project/ZapDemo/log/test.log",
		MaxSize:    10, //MB
		MaxBackups: 2,
		MaxAge:     28,   // days
		Compress:   true, // 是否压缩文件
	}
	// 初始化logger
	zapLogger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(lumberjackLogger),
		zap.InfoLevel,
	))
	defer zapLogger.Sync()

	for i := 0; i < 100; i++ {
		str := fmt.Sprintf("hi %d", i)
		zapLogger.Info(str)
	}

}
