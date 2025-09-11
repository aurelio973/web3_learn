package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/natefinch/lumberjack"
)

// 全局日志实例，供其他包调用
var Logger *zap.Logger

// InitLogger 初始化日志配置
func InitLogger() {
	// 日志文件配置（切割、备份等）
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/blog.log", // 日志文件路径（项目根目录下的logs文件夹）
		MaxSize:    100,              // 单个文件最大100MB
		MaxBackups: 10,               // 最多保留10个备份文件
		MaxAge:     30,               // 保留30天
		Compress:   true,             // 压缩旧日志
	}

	// 日志输出格式配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式：2024-05-20T15:04:05Z07:00
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志级别大写（INFO/ERROR等）

	// 构建日志核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // JSON格式输出
		zapcore.AddSync(lumberJackLogger),     // 输出到文件
		zapcore.DebugLevel,                    // 最低日志级别（Debug及以上都记录）
	)

	// 创建日志实例
	Logger = zap.New(core, zap.AddCaller()) // 添加调用者信息（哪个文件哪一行）

	// 测试日志输出
	Logger.Info("日志初始化成功")
}
    
