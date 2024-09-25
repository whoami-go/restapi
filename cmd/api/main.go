package main

import (
	"awesomeProject/internal/app/api"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string = "configs/config.toml"
)

func init() {
	//скажем что наше приложение будет получать на этапе запуска путь к конфиг файлу из вне
	flag.StringVar(&configPath, "path", "configs/config.toml", "path to config file .toml format")
}
func main() {
	//в этот момент происходит инициализация переменной конфиг пас значением
	flag.Parse()
	log.Println("its work correct")

	//server instance initialization !!
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	// дессириализуем содержимое .toml
	if err != nil {
		log.Fatal("can not find configs file, using default values:", err)
	}

	//теперь тут надо попробовать прочитать из .toml так как там может быть новая инфо
	server := api.New(config)

	//server api start
	log.Fatal(server.Start())
}
