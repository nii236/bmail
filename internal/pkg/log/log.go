package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var z *zap.SugaredLogger

func NewToFile(filename string) *zap.SugaredLogger {
	conf := zap.NewDevelopmentConfig()
	conf.OutputPaths = []string{
		filename,
	}
	l, err := conf.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.FatalLevel),
	)
	if err != nil {
		panic(err)
	}
	z = l.Sugar()
	return z
}
func New(production bool) *zap.SugaredLogger {
	conf := zap.NewDevelopmentConfig()
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := conf.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.FatalLevel),
	)

	if production {
		l, err = zap.NewProduction(
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zap.FatalLevel),
		)
	}

	if err != nil {
		panic(err)
	}

	z = l.Sugar()
	return z
}

// Get returns the package global
func Get() *zap.SugaredLogger {
	if z == nil {
		panic("log not initialized")
	}
	return z
}

// Fatal is a wrapper for zap's Fatal
func Fatal(args ...interface{}) {
	z.Fatal(args...)
}

// Print is a wrapper for zap's Print
func Print(args ...interface{}) {
	z.Info(args...)
}

// Info is a wrapper for zap's Info
func Info(args ...interface{}) {
	z.Info(args...)
}

// Warn is a wrapper for zap's Warn
func Warn(args ...interface{}) {
	z.Warn(args...)
}

// Error is a wrapper for zap's Error
func Error(args ...interface{}) {
	z.Error(args...)
}

// Debug is a wrapper for zap's Debug
func Debug(args ...interface{}) {
	z.Debug(args...)
}

// Fatalln is a wrapper for zap's Fatalln
func Fatalln(args ...interface{}) {
	z.Fatal(args...)
}

// Println is a wrapper for zap's Println
func Println(args ...interface{}) {
	z.Info(args...)
}

// Infoln is a wrapper for zap's Infoln
func Infoln(args ...interface{}) {
	z.Info(args...)
}

// Warnln is a wrapper for zap's Warnln
func Warnln(args ...interface{}) {
	z.Warn(args...)
}

// Errorln is a wrapper for zap's Errorln
func Errorln(args ...interface{}) {
	z.Error(args...)
}

// Debugln is a wrapper for zap's Debugln
func Debugln(args ...interface{}) {
	z.Debug(args...)
}

// Fatalw is a wrapper for zap's Fatalw
func Fatalw(msg string, keysandvalues ...interface{}) {
	z.Fatalw(msg, keysandvalues...)
}

// Printw is a wrapper for zap's Printw
func Printw(msg string, keysandvalues ...interface{}) {
	z.Infow(msg, keysandvalues...)
}

// Infow is a wrapper for zap's Infow
func Infow(msg string, keysandvalues ...interface{}) {
	z.Infow(msg, keysandvalues...)
}

// Warnw is a wrapper for zap's Warnw
func Warnw(msg string, keysandvalues ...interface{}) {
	z.Warnw(msg, keysandvalues...)
}

// Errorw is a wrapper for zap's Errorw
func Errorw(msg string, keysandvalues ...interface{}) {
	z.Errorw(msg, keysandvalues...)
}

// Debugw is a wrapper for zap's Debugw
func Debugw(msg string, keysandvalues ...interface{}) {
	z.Debugw(msg, keysandvalues...)
}

// Printf is a wrapper for zap's Printf
func Printf(format string, args ...interface{}) {
	z.Infof(format, args...)
}

// Infof is a wrapper for zap's Infof
func Infof(format string, args ...interface{}) {
	z.Infof(format, args...)
}

// Warnf is a wrapper for zap's Warnf
func Warnf(format string, args ...interface{}) {
	z.Warnf(format, args...)
}

// Errorf is a wrapper for zap's Errorf
func Errorf(format string, args ...interface{}) {
	z.Errorf(format, args...)
}

// Debugf is a wrapper for zap's Debugf
func Debugf(format string, args ...interface{}) {
	z.Debugf(format, args...)
}

// Fields wraps logrus fields
type Fields map[string]interface{}

// WithFields is a wrapper for zap's WithFields
func WithFields(fields Fields) *Entry {
	// entry := logrus.WithFields(logrus.Fields(fields))
	for k, v := range fields {
		z.With(k, v)
	}
	return &Entry{z}
}

// Entry wraps the logrus Entry type
type Entry struct {
	*zap.SugaredLogger
}

// WithField is a wrapper for zap's WithField
func WithField(key string, value interface{}) *Entry {
	z.With(key, value)
	return &Entry{z}
}
