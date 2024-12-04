package api

import (
	"awesomeProject/internal/app/middleware"
	"awesomeProject/internal/app/models"
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

// вспомогательная структура для сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

// full API handlers initializations
func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// возвращает все статьи из бд на данный момент
func (api *API) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	//инициализируем хедеры
	initHeaders(w)
	//логируем момент начала обработки запроса
	api.logger.Info("Get All Articles GET /api/v1/articles")
	//пытаемся что то получить от бд
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		//что делаем если была ошибка на этапе подключения
		api.logger.Info("Error while  Articles.SelectAll : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles try again later.",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(articles)
}

func (api *API) GetArticleById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("get by id")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Troubles while id param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Dont use ID as casting to int value",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Trouble while accessing db table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing db",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not article with that ID in db")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exist in db",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(article)
}

func (api *API) DeleteByID(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Delete Article by ID")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	//удостоверились, что передали валидный параметр и его можно привести к целому числу
	if err != nil {
		api.logger.Info("Troubles while id param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Dont use ID as casting to int value",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	//удостоверимся что соединение в БД еще есть и запросы формируются
	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Trouble while accessing db table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing db",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// проверяем то что мы хотим удалить находится в БД
	if !ok {
		api.logger.Info("Can not article with that ID in db")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exist in db",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Trouble while deleting DB element with ID")
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trouble to accessing DB",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Artical with ID %d successfully deleted", id),
		IsError:    false,
	}
	json.NewEncoder(w).Encode(msg)

}

func (api *API) PostArticle(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid json received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provide json is invalid",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating article :", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing db. Try again.",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)

	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(a)
}

func (api *API) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Post User register request")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provide json is invalid",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	//пытаемся найти пользователя с таким логином БД
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Trouble while accessing db table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing db",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	//если такой пользователь уже есть - то регистрации не будет !
	if !ok {
		api.logger.Info("User with this ID already exist")
		msg := Message{
			StatusCode: 400,
			Message:    "User with this ID already exist in DB ",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	//теперь пытаемся добавить в БД
	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Trouble while accessing db table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing db",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} sucessfully registrtation! ", userAdded.Login),
		IsError:    false,
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(msg)

}

func (api *API) PostToAuth(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Post To Auth POST /api/v1/auth")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	//обрабатываем случай если json вовсе не json или в нем какие либо проблемы
	if err != nil {
		api.logger.Info("Invalid json received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provide json is invalid",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	//необходимо попытаться обнаружить с таким логин в БД
	userInDB, ok, err := api.storage.User().FindByLogin(user.Login)
	//проблема доступа к бд
	if err != nil {
		api.logger.Info("Trouble while accessing db table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing db",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	//если подключение удалось но пользователя с таким логином нет
	if !ok {
		api.logger.Info("User with that login does not exist:", err)
		msg := Message{
			StatusCode: 404,
			Message:    "User with that ID does not exist in db",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	// Если юзер есть в бд с таким логином сравним пароли
	if userInDB.Password != user.Password {
		api.logger.Info("User with that login does not match password")
		msg := Message{
			StatusCode: 404,
			Message:    "User with that login does not match password",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return

	}

	// теперь выбиваем токен как знак успешной аутентификации
	token := jwt.New(jwt.SigningMethodHS256)             // тот же метод подписания токена как и в JwtMiddleware.go
	claims := token.Claims.(jwt.MapClaims)               // дополнительные дейсвтия (в формате мапы)  для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Ограничение 2 часа на пользование ( время жизни токена)
	claims["admin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey) //метомд который будет генерировать строку JWT

	// В случае если токен выбить не удалось по каким либо причинам
	if err != nil {
		api.logger.Info("Can not claim jwt - token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing db",
			IsError:    true,
		}
		// в случае если токен видим и отдаем его клиенту
		msg = Message{
			StatusCode: 201,
			Message:    tokenString,
			IsError:    false,
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(msg)
	}

}
