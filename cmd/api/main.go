package main

import (
	"awesomeProject/internal/app/api"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	//скажем что наше приложение будет получать на этапе запуска путь к конфиг файлу из вне
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file .toml format")
}
func main() {
	//в этот момент происходит инициализация переменной конфиг пас значением
	flag.Parse()
	log.Fatal("its work")

	//server instance initialization
	config := api.NewConfig()
	_, err := toml.Decode(configPath, config) // дессириализуем содержимое .toml
	if err != nil {
		log.Fatal("can not find configs file, using default values:", err)
	}
	//теперь тут надо попробовать прочитать из .toml так как там может быть новая инфо
	server := api.New(config)

	//server api start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
