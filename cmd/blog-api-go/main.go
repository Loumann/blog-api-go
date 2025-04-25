package main

import (
	"blog-api-go/config"
	"blog-api-go/internal/controller"
	"blog-api-go/internal/models"
	repos2 "blog-api-go/internal/repos"
	"blog-api-go/internal/service"
	"fmt"
	"log"
)

func main() {
	cfg := config.ReadCfg()
	envs := models.LoadEnv()

	fmt.Println("envs: ", envs, cfg)

	db := repos2.NewBusinessDatabase(cfg, envs)

	Repos := repos2.NewRepository(db)
	Service := service.NewService(Repos)
	Controller := controller.NewController(Service)

	router := Controller.InitRouters()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Fatal::", err)
	}

	return

}
