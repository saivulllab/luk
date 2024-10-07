package logger

import (
	"enum"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var zapLogger *zap.Logger

// Environment represents the type of environment
type Environment string

// Ensure Environment implements enum.EnumValid interface
var _ enum.EnumValid = (*Environment)(nil)

// Ensure Environment implements enum.EnumStringable interface
var _ enum.EnumStringable = (*Environment)(nil)

const (
	Production Environment = "prod"
	Local      Environment = "local"
)

func (e Environment) Valid() bool {
	switch e {
	case Production, Local:
		return true
	}

	return false
}

func (e *Environment) FromString(str string) error {
	val := Environment(str)

	if !val.Valid() {
		return fmt.Errorf("wrong enum value: %s", val)
	}

	*e = val

	return nil
}

func (e Environment) ToString() string {
	return string(e)
}

func lowerCaseLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	if level == zap.PanicLevel || level == zap.DPanicLevel {
		enc.AppendString("error")
		return
	}
	zapcore.LowercaseLevelEncoder(level, enc)
}

func NewLogger(env string) *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	var level zapcore.Level
	var encoding string
	var sampling bool

	switch Environment(env) {
	case Production:
		level = zap.InfoLevel
		encoding = "json"
		sampling = true
	case Local:
		level = zap.DebugLevel
		encoding = "console"
		sampling = false
	default:
		level = zap.InfoLevel
		encoding = "json"
		sampling = true
	}

	atomicLevel := zap.NewAtomicLevelAt(level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    lowerCaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder

	if encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, stdout, atomicLevel)

	if sampling {
		core = zapcore.NewSamplerWithOptions(
			core,
			time.Second,
			3,
			0,
		)
	}

	zapLogger = zap.New(core, zap.AddCaller())

	return zapLogger
}

func Get() *zap.Logger {
	return zapLogger
}
