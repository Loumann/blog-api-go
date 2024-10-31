package main

import (
	"blog-api-go/config"
	"blog-api-go/controller"
	"blog-api-go/models"
	"blog-api-go/repos"
	"blog-api-go/service"
	"fmt"
	"log"
)

func main() {
	cfg := config.ReadCfg()
	envs := models.LoadEnv()

	fmt.Println("envs: ", envs, cfg)

	db := repos.NewBusinessDatabase(cfg, envs)

	Repos := repos.NewRepository(db)
	Service := service.NewService(Repos)
	Controller := controller.NewController(Service)

	router := Controller.InitRouters()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Fatal::", err)
	}
	return

}
