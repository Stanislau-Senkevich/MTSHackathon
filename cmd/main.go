package main

import (
	"MTSHackathonBackEnd/internal/config"
	"MTSHackathonBackEnd/internal/handler"
	"MTSHackathonBackEnd/internal/repository"
	mongodb "MTSHackathonBackEnd/internal/repository/mongo"
	"MTSHackathonBackEnd/internal/server"
	"MTSHackathonBackEnd/internal/service"
	"fmt"
	"log"
	"sync"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error due reading config: %s", err.Error())
	}
	mongo, err := mongodb.InitMongoRepository(&cfg.Mongo)
	if err != nil {
		log.Fatalf(err.Error())
	}

	srv := new(server.Server)
	repos := repository.NewRepository(mongo)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err = srv.Run(cfg.Port, handlers.InitRoutes())
		if err != nil {
			log.Fatalf("error due running server: %s", err.Error())
		}
	}()

	fmt.Print("\n\n\n\n\n\n")
	fmt.Println(cfg)
	fmt.Print("\n\n\n\n\n\n")

	wg.Wait()
}
