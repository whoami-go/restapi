package api

import "github.com/sirupsen/logrus"

// пытаемся отконфигурировать наш API instance (logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil

}
