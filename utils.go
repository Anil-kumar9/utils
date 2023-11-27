package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type VinLogger struct {
	*logrus.Logger
	VinEnabled bool
}

func NewLogger(vinEnabled bool) VinLogger {
	logger := logrus.New()
	VinLogger := VinLogger{Logger: logger, VinEnabled: vinEnabled}
	return VinLogger
}

func (logger *VinLogger) Info(msg string) {
	if isLog(logger.VinEnabled) {
		logrus.Info(msg)
	}
}
func (logger *VinLogger) Warn(msg string) {
	if isLog(logger.VinEnabled) {
		logrus.Warn(msg)
	}
}
func (logger *VinLogger) Error(msg string) {
	logrus.Error(msg)
}

func isLog(vinEnabled bool) bool {
	if viper.GetBool("Log-Enabled") || vinEnabled {
		return true
	}
	return false
}
