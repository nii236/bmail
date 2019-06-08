package log

import (
	"go.uber.org/zap"
)

// Logger is a single logger instance
type Logger struct {
	*zap.SugaredLogger
}

// New creates a new logger that writes to a file
func New() *Logger {
	conf := zap.NewDevelopmentConfig()
	l, err := conf.Build(
		zap.AddStacktrace(zap.FatalLevel),
	)
	if err != nil {
		panic(err)
	}
	z := l.Sugar()
	return &Logger{z}
}

// NewToFile creates a new logger that writes to a file
func NewToFile(filename string) *Logger {
	conf := zap.NewDevelopmentConfig()
	conf.OutputPaths = []string{
		filename,
	}
	l, err := conf.Build(
		zap.AddStacktrace(zap.FatalLevel),
	)
	if err != nil {
		panic(err)
	}
	z := l.Sugar()
	return &Logger{z}
}

// Fatal is a wrapper for zap's Fatal
func (z *Logger) Fatal(args ...interface{}) {
	z.SugaredLogger.Fatal(args...)
}

// Print is a wrapper for zap's Print
func (z *Logger) Print(args ...interface{}) {
	z.SugaredLogger.Info(args...)
}

// Info is a wrapper for zap's Info
func (z *Logger) Info(args ...interface{}) {
	z.SugaredLogger.Info(args...)
}

// Warn is a wrapper for zap's Warn
func (z *Logger) Warn(args ...interface{}) {
	z.SugaredLogger.Warn(args...)
}

// Error is a wrapper for zap's Error
func (z *Logger) Error(args ...interface{}) {
	z.SugaredLogger.Error(args...)
}

// Debug is a wrapper for zap's Debug
func (z *Logger) Debug(args ...interface{}) {
	z.SugaredLogger.Debug(args...)
}

// Fatalln is a wrapper for zap's Fatalln
func (z *Logger) Fatalln(args ...interface{}) {
	z.SugaredLogger.Fatal(args...)
}

// Println is a wrapper for zap's Println
func (z *Logger) Println(args ...interface{}) {
	z.SugaredLogger.Info(args...)
}

// Infoln is a wrapper for zap's Infoln
func (z *Logger) Infoln(args ...interface{}) {
	z.SugaredLogger.Info(args...)
}

// Warnln is a wrapper for zap's Warnln
func (z *Logger) Warnln(args ...interface{}) {
	z.SugaredLogger.Warn(args...)
}

// Errorln is a wrapper for zap's Errorln
func (z *Logger) Errorln(args ...interface{}) {
	z.SugaredLogger.Error(args...)
}

// Debugln is a wrapper for zap's Debugln
func (z *Logger) Debugln(args ...interface{}) {
	z.SugaredLogger.Debug(args...)
}

// Fatalw is a wrapper for zap's Fatalw
func (z *Logger) Fatalw(msg string, keysandvalues ...interface{}) {
	z.SugaredLogger.Fatalw(msg, keysandvalues...)
}

// Printw is a wrapper for zap's Printw
func (z *Logger) Printw(msg string, keysandvalues ...interface{}) {
	z.SugaredLogger.Infow(msg, keysandvalues...)
}

// Infow is a wrapper for zap's Infow
func (z *Logger) Infow(msg string, keysandvalues ...interface{}) {
	z.SugaredLogger.Infow(msg, keysandvalues...)
}

// Warnw is a wrapper for zap's Warnw
func (z *Logger) Warnw(msg string, keysandvalues ...interface{}) {
	z.SugaredLogger.Warnw(msg, keysandvalues...)
}

// Errorw is a wrapper for zap's Errorw
func (z *Logger) Errorw(msg string, keysandvalues ...interface{}) {
	z.SugaredLogger.Errorw(msg, keysandvalues...)
}

// Debugw is a wrapper for zap's Debugw
func (z *Logger) Debugw(msg string, keysandvalues ...interface{}) {
	z.SugaredLogger.Debugw(msg, keysandvalues...)
}

// Printf is a wrapper for zap's Printf
func (z *Logger) Printf(format string, args ...interface{}) {
	z.SugaredLogger.Infof(format, args...)
}

// Infof is a wrapper for zap's Infof
func (z *Logger) Infof(format string, args ...interface{}) {
	z.SugaredLogger.Infof(format, args...)
}

// Warnf is a wrapper for zap's Warnf
func (z *Logger) Warnf(format string, args ...interface{}) {
	z.SugaredLogger.Warnf(format, args...)
}

// Errorf is a wrapper for zap's Errorf
func (z *Logger) Errorf(format string, args ...interface{}) {
	z.SugaredLogger.Errorf(format, args...)
}

// Debugf is a wrapper for zap's Debugf
func (z *Logger) Debugf(format string, args ...interface{}) {
	z.SugaredLogger.Debugf(format, args...)
}
