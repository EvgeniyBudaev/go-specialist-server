package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/EvgeniyBudaev/go-specialist-server/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	// Наше приложение на этапе запуска будет получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse() // Инициализация переменной configPath значениями из flag.StringVar()
	log.Println("It's work")
	// Server instance initialization
	config := api.NewConfig()
	_, err := toml.Decode(configPath, config) // Десериализация содержимого .toml файла
	//fmt.Println(res)
	if err != nil {
		log.Println("can't find configs file. Using default values:", err)
	}

	server := api.New(config)

	// API server start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	} // или в одну строчку log.Fatal(server.Start())
}
