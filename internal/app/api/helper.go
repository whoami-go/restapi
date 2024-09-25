package api

import (
	"awesomeProject/storage"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// пытаемся отконфигурировать наш API instance (logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil

}

// пытаемся отконфигурировать маршрутизатор (роутер апи )
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ! This is rest api!"))
	})
}

// пытаемся конфигурировать наше хранилище (storage API)
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//пытаемся установить соединение если невозможно - return err
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
