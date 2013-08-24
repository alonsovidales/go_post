package logger

import (
    "config"
    "log"
)

var logLevel = config.Get().Default.Logging

func Debug(format string, v ...interface{}) {
    if logLevel == "DEBUG" {
        log.Printf(format, v)
    }
}

func Info(format string, v ...interface{}) {
    if logLevel == "DEBUG" || logLevel == "INFO" {
        log.Printf(format, v)
    }
}

func Panic(format string, v ...interface{}) {
    if logLevel == "DEBUG" || logLevel == "INFO" || logLevel == "PANIC" {
        log.Panicf(format, v)
    }
}

func Fatal(format string, v ...interface{}) {
    log.Fatalf(format, v)
}
