package main

import "fmt"

type OldLogger struct{}

func (ll *OldLogger) WriteLog(level int, message string) {
	levels := map[int]string{
		1: "DEBUG",
		2: "INFO",
		3: "WARN",
		4: "ERROR",
	}
	fmt.Printf("[%s] %s\n", levels[level], message)
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type LoggerAdapter struct {
	OldLogger *OldLogger
}

func NewLoggerAdapter(OldLogger *OldLogger) *LoggerAdapter {
	return &LoggerAdapter{OldLogger: OldLogger}
}

func (la *LoggerAdapter) Debug(msg string) {
	la.OldLogger.WriteLog(1, msg)
}

func (la *LoggerAdapter) Info(msg string) {
	la.OldLogger.WriteLog(2, msg)
}

func (la *LoggerAdapter) Warn(msg string) {
	la.OldLogger.WriteLog(3, msg)
}

func (la *LoggerAdapter) Error(msg string) {
	la.OldLogger.WriteLog(4, msg)
}

type ModernApplication struct {
	logger Logger
}

func NewModernApplication(logger Logger) *ModernApplication {
	return &ModernApplication{logger: logger}
}

func (app *ModernApplication) ProcessData() {
	app.logger.Info("Application start")
	app.logger.Debug("Settings: some settings")
	app.logger.Warn("Warning: something  went wrong")
	app.logger.Error("Error: some error occurred")
}

func main() {
	OldLogger := &OldLogger{}

	loggerAdapter := NewLoggerAdapter(OldLogger)

	app := NewModernApplication(loggerAdapter)

	app.ProcessData()
}
