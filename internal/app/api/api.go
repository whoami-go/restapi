package api

import (
	"awesomeProject/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// base api server instance description
type API struct {
	//unexported field!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//добавление поля для работы с хранилищем
	storage *storage.Storage
}

// api constructor : build base api instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
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
	//конфигурируем маршрутизатор
	api.configureRouterField()
	//конфигурируем хранилище
	if err := api.configureStorageField(); err != nil {
		return err
	}
	//на этапе валидного завершения старнуем http server
	return http.ListenAndServe(api.config.BindAddr, api.router)

}
