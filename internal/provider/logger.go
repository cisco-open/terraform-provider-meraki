package provider

import (
	"log"
	"os"
)

func createLogger() *logger {
	l := &logger{l: log.New(os.Stderr, "", log.LstdFlags)}
	return l
}

type logger struct {
	l *log.Logger
}

func (l *logger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] [RESTY] "+format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	log.Printf("[WARN] [RESTY] "+format, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG] [RESTY] "+format, v...)
}
