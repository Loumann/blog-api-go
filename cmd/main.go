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

	userRepo := repos.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := userController.InitRouters()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Fatal::", err)
	}
	return

}
