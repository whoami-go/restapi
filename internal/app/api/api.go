package api

import (
	"github.com/sirupsen/logrus"
)

// base api server instance description
type API struct {
	//unexported field!
	config *Config
	logger *logrus.Logger
}

// api constructor : build base api instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

// start http server / configure logger , router and etc
func (api *API) Start() error {
	//trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//подтверждение того что логгер сконфигурировался
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	return nil
}
