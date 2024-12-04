package api

import (
	"awesomeProject/storage"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
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

// пытаемся отконфигурировать маршрутизатор (router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteByID).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")

	// new pair for auth
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")

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
