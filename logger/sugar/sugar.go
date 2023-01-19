package sugar

import (
	"github.com/halilylm/utils/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sugarLogger struct {
	level           logger.Level
	developmentMode bool
	sugarLogger     *zap.SugaredLogger
	initialFields   map[string]any
}

type Options struct {
	Level           logger.Level
	DevelopmentMode bool
	InitialFields   map[string]any
}

func NewApiLogger(opts *Options) (logger.Logger, error) {
	s := &sugarLogger{
		level:           opts.Level,
		developmentMode: opts.DevelopmentMode,
		initialFields:   opts.InitialFields,
	}
	logLevel := s.mapLogLevel()
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: s.developmentMode,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "ts",
			NameKey:       "name",
			CallerKey:     "caller",
			FunctionKey:   "function",
			StacktraceKey: "stack",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
		InitialFields: s.initialFields,
	}
	l, err := logConfig.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	s.sugarLogger = l.Sugar()
	return s, nil
}

// mapLogLevels includes related log levels
var mapLogLevels = map[logger.Level]zapcore.Level{
	logger.DebugLevel: zapcore.DebugLevel,
	logger.InfoLevel:  zapcore.InfoLevel,
	logger.WarnLevel:  zapcore.WarnLevel,
	logger.ErrorLevel: zapcore.ErrorLevel,
	logger.PanicLevel: zapcore.PanicLevel,
	logger.FatalLevel: zapcore.FatalLevel,
}

// mapLogLevel matches log level with zap package
func (s *sugarLogger) mapLogLevel() zapcore.Level {
	if level, ok := mapLogLevels[s.level]; ok {
		return level
	}
	return zapcore.DebugLevel
}

func (s *sugarLogger) Sync() error {
	return s.sugarLogger.Sync()
}

func (s *sugarLogger) Debug(args ...interface{}) {
	s.sugarLogger.Debug(args...)
}

func (s *sugarLogger) Debugf(template string, args ...interface{}) {
	s.sugarLogger.Debugf(template, args...)
}

func (s *sugarLogger) Info(args ...interface{}) {
	s.sugarLogger.Info(args...)
}

func (s *sugarLogger) Infof(template string, args ...interface{}) {
	s.sugarLogger.Infof(template, args...)
}

func (s *sugarLogger) Warn(args ...interface{}) {
	s.sugarLogger.Warn(args...)
}

func (s *sugarLogger) Warnf(template string, args ...interface{}) {
	s.sugarLogger.Warnf(template, args...)
}

func (s *sugarLogger) Error(args ...interface{}) {
	s.sugarLogger.Error(args...)
}

func (s *sugarLogger) Errorf(template string, args ...interface{}) {
	s.sugarLogger.Errorf(template, args...)
}

func (s *sugarLogger) Panic(args ...interface{}) {
	s.sugarLogger.Panic(args...)
}

func (s *sugarLogger) Panicf(template string, args ...interface{}) {
	s.sugarLogger.Panicf(template, args...)
}

func (s *sugarLogger) Fatal(args ...interface{}) {
	s.sugarLogger.Fatal(args...)
}

func (s *sugarLogger) Fatalf(template string, args ...interface{}) {
	s.sugarLogger.Fatalf(template, args...)
}
