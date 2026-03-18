package global

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
)

func NewLogger() (*zap.Logger, error) {
	cfg := GetLoggerConfig()
	level := zap.NewAtomicLevel()
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	var base zap.Config
	if cfg.Development {
		base = zap.NewDevelopmentConfig()
	} else {
		base = zap.NewProductionConfig()
	}

	base.Encoding = cfg.Encoding
	base.Level = level
	base.EncoderConfig.TimeKey = "timestamp"
	base.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 如果配置了日志文件，则使用自定义核心支持日志轮转
	if cfg.File != "" {
		logger, err := buildLoggerWithRotation(cfg, base, level)
		if err != nil {
			return nil, err
		}
		zap.ReplaceGlobals(logger)
		return logger, nil
	}

	// 否则使用默认配置
	if err := applyOutput(&base, cfg.File); err != nil {
		return nil, err
	}

	logger, err := base.Build()
	if err != nil {
		return nil, fmt.Errorf("build logger: %w", err)
	}

	zap.ReplaceGlobals(logger)
	return logger, nil
}

func buildLoggerWithRotation(cfg LoggerConfig, base zap.Config, level zap.AtomicLevel) (*zap.Logger, error) {
	// 创建日志目录
	dir := filepath.Dir(cfg.File)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, fmt.Errorf("create log dir: %w", err)
		}
	}

	// 使用lumberjack进行日志轮转
	lumberjackWriter := &lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	}

	// 创建编码器
	encoderConfig := base.EncoderConfig
	var encoder zapcore.Encoder
	if base.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 解析日志级别
	logLevel := level.Level()

	// 创建核心
	var cores []zapcore.Core

	// 添加标准输出核心
	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), logLevel))

	// 添加文件输出核心（带轮转）
	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(lumberjackWriter), logLevel))

	// 创建logger
	options := []zap.Option{zap.AddCaller()}
	if !cfg.Development {
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	}

	logger := zap.New(zapcore.NewTee(cores...), options...)
	return logger, nil
}

func applyOutput(base *zap.Config, filePath string) error {
	outputPaths := []string{"stdout"}
	errorOutputPaths := []string{"stderr"}

	if filePath != "" {
		dir := filepath.Dir(filePath)
		if dir != "." && dir != "" {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("create log dir: %w", err)
			}
		}
		outputPaths = append(outputPaths, filePath)
		errorOutputPaths = append(errorOutputPaths, filePath)
	}

	base.OutputPaths = outputPaths
	base.ErrorOutputPaths = errorOutputPaths
	return nil
}
