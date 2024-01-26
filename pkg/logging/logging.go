package logging

import (
	"fmt"
	"log/slog"
	"os"
	"path"
)

type Logging interface {
	Debug(mesg string, args ...any)
	Info(mesg string, args ...any)
	Warn(mesg string, args ...any)
	Error(mesg string, args ...any)
	Fatal(mesg string, args ...any)
	Debugf(mesg string, args ...any)
	Infof(mesg string, args ...any)
	Warnf(mesg string, args ...any)
	Errorf(mesg string, args ...any)
	Fatalf(mesg string, args ...any)
}

type Logger struct {
	*slog.Logger
}

func (lgr *Logger) Debugf(mesg string, args ...any) {
	lgr.Logger.Debug(fmt.Sprintf(mesg, args...))
}

func (lgr *Logger) Infof(mesg string, args ...any) {
	lgr.Logger.Info(fmt.Sprintf(mesg, args...))
}

func (lgr *Logger) Warnf(mesg string, args ...any) {
	lgr.Logger.Warn(fmt.Sprintf(mesg, args...))
}

func (lgr *Logger) Errorf(mesg string, args ...any) {
	lgr.Logger.Error(fmt.Sprintf(mesg, args...))
}

func (lgr *Logger) Fatal(mesg string, args ...any) {
	lgr.Logger.Error(mesg, args...)
	os.Exit(1)
}

func (lgr *Logger) Fatalf(mesg string, args ...any) {
	lgr.Logger.Error(fmt.Sprintf(mesg, args...))
	os.Exit(1)
}

func New(level slog.Level, isConsole bool, paths ...string) Logging {
	fPath := path.Join(paths...)
	if fPath == "" {
		fPath = "log.log"
	}
	var writer *os.File
	var err error
	if !isConsole {
		writer, err = os.OpenFile(fPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			slog.Debug("Can't open log file")
			writer = os.Stdout
		}
	} else {
		writer = os.Stdout
	}
	logger := slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	return &Logger{
		Logger: logger,
	}
}
