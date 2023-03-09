package main

import (
	"go-blogrpl/config"
	"go-blogrpl/controller"
	"go-blogrpl/middleware"
	"go-blogrpl/repository"
	"go-blogrpl/routes"
	"go-blogrpl/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// API Documentation : https://documenter.getpostman.com/view/25087235/2s93CUJVem

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	// Setting Up Database
	db := config.DBSetup()

	// Setting Up Repositories
	userR := repository.NewUserRepository(db)

	// Setting Up Services
	userS := service.NewUserService(userR)
	jwtS := service.NewJWTService()

	// Setting Up Controllers
	userC := controller.NewUserController(userS, jwtS)

	defer config.DBClose(db)

	// Setting Up Server
	server := gin.Default()
	server.Use(
		middleware.CORSMiddleware(),
	)

	// Setting Up Routes
	routes.UserRoutes(server, userC)

	// Running in localhost:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
