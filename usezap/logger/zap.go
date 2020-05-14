package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *ZapLogger

type ZapLogger struct {
	*zap.SugaredLogger
	logLevel zap.AtomicLevel
}

type MyLoginfo struct {
	LogInfo
	Name string
}

type LogInfo struct {
	Level        string `json:"level"`          // 日志等级
	OutputPath   string `json:"output_path"`    // 输出地址
	MaxSizeMb    int    `json:"max_size_mb"`    // 日志文件大小，单位mb，默认100
	MaxBackups   int    `json:"max_backups"`    // 最大备份数量，默认10
	JsonFormat   bool   `json:"json_format"`    // 是否json输出
	LogInConsole bool   `json:"log_in_console"` // 是否同时输出到console
}

const defaultloglevel = zapcore.DebugLevel // 默认日志等级为debug

//设置日志
// Param:
// logPath 日志文件路径
// logLevel 日志级别 debug/info/warn/error
// maxSize 单个文件大小,MB
// maxBackups 保存的文件个数
// compress 压缩
// jsonFormat 是否输出为json格式
// shoowLine 显示代码行
// logInConsole 是否同时输出到控制台
func (l *ZapLogger) setLogger(logLevel, logPath string, maxSizeMb, maxBackups int, jsonFormat, logInConsole bool) {
	hook := lumberjack.Logger{
		Filename:   logPath,    // 日志文件路径
		MaxSize:    maxSizeMb,  // megabytes
		MaxBackups: maxBackups, // 最多保留300个备份
		//Compress:   false,      // 是否压缩 disabled by default
	}

	var syncer zapcore.WriteSyncer
	if logInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		syncer = zapcore.AddSync(&hook)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	if jsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	if level.UnmarshalText([]byte(logLevel)) != nil {
		level = defaultloglevel
	}

	atomicLevel := zap.NewAtomicLevelAt(level)

	core := zapcore.NewCore(
		encoder,
		syncer,
		atomicLevel,
	)

	logger := zap.New(core)
	logger = logger.WithOptions(zap.AddCaller()) //添加代码行号
	l.SugaredLogger = logger.Sugar()
	l.logLevel = atomicLevel
}

func InitLogger(logConfig *LogInfo) error {
	if logConfig == nil {
		return fmt.Errorf("logConfig is null")
	}

	Log = &ZapLogger{}

	var level zapcore.Level
	if level.UnmarshalText([]byte(logConfig.Level)) != nil {
		level = zapcore.DebugLevel
	}

	Log.setLogger(logConfig.Level, logConfig.OutputPath, logConfig.MaxSizeMb, logConfig.MaxBackups, logConfig.JsonFormat, logConfig.LogInConsole)
	if Log.SugaredLogger == nil {
		return fmt.Errorf("log SugaredLogger is null")
	}
	return nil
}

func (l *ZapLogger) SetLevel(logLevel string) error {
	var level zapcore.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return err
	}

	l.logLevel.SetLevel(level)
	return nil
}

func (l *ZapLogger) Close() {
	l.SugaredLogger.Sync()
}
