package logger

import (
	"github.com/sirupsen/logrus"
)

// logrusLogger for logging using logrus
type logrusLogger struct {
	logger *logrus.Logger
	hooks  []logrus.Hook
	level  string
}

// NewLogrusLogger is the factory function for logrus logger
func NewLogrusLogger(level string, hooks []logrus.Hook) Logger {
	return &logrusLogger{level: level, hooks: hooks}
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
}

func (l *logrusLogger) getLoggerLevel() logrus.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return logrus.DebugLevel
	}
	return level
}

// InitLogger set ups the logging library
func (l *logrusLogger) InitLogger(reqFields RequiredFields) {
	logLevel := l.getLoggerLevel()
	logger := &logrus.Logger{}
	logger.SetLevel(logLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	for _, hook := range l.hooks {
		logger.AddHook(hook)
	}
	logger.WithField("release", reqFields.Release)
	l.logger = logger
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logrusLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}
