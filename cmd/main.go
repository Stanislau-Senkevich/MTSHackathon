package main

import (
	"MTSHackathonBackEnd/internal/config"
	"MTSHackathonBackEnd/internal/handler"
	"MTSHackathonBackEnd/internal/repository"
	mongodb "MTSHackathonBackEnd/internal/repository/mongo"
	"MTSHackathonBackEnd/internal/server"
	"MTSHackathonBackEnd/internal/service"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error due reading config: %s", err.Error())
	}
	mongo, err := mongodb.InitMongoRepository(&cfg.Mongo)
	defer func() {
		if err = mongo.DB.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	srv := new(server.Server)
	repos := repository.NewRepository(mongo)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go func() {
		err = srv.Run(cfg.Port, handlers.InitRoutes())
		if err != nil {
			log.Fatalf("error due running server: %s", err.Error())
		}
	}()

	time.Sleep(time.Second * 2)
	cat, _ := repos.GetAllCategories()
	fmt.Println(cat)

}
