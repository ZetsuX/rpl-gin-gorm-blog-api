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

	// Setting Up Controllers
	userC := controller.NewUserController(userS)

	defer config.DBClose(db)

	// Setting Up Server
	server := gin.Default()
	server.Use(
		middleware.CORSMiddleware(),
	)

	// Setting Up Routes
	routes.UserRoutes(server, userC)

	// // Handlers
	// nasabahHandler := handler.NasabahHandler{DB: db}
	// noTelpHandler := handler.NoTelpHandler{DB: db}
	// rekeningHandler := handler.RekeningHandler{DB: db}

	// // Nasabah Routes
	// server.GET("/nasabah", nasabahHandler.HandleGetNasabah)
	// server.GET("/nasabah/:id", nasabahHandler.HandleGetNasabahByID)
	// server.POST("/nasabah", nasabahHandler.HandleInsertNasabah)
	// server.PUT("/nasabah/:id", nasabahHandler.HandleEditNasabahByID)
	// server.DELETE("/nasabah/:id", nasabahHandler.HandleDeleteNasabahByID)

	// // No Telp Routes
	// server.GET("/notelp/nasabah/:id", noTelpHandler.HandleGetNoTelpByNasabahID)
	// server.GET("/notelp", noTelpHandler.HandleGetNoTelp)
	// server.GET("/notelp/:id", noTelpHandler.HandleGetNoTelpByID)
	// server.POST("/notelp", noTelpHandler.HandleInsertNoTelp)
	// server.DELETE("/notelp/:id", noTelpHandler.HandleDeleteNoTelpByID)
	// server.PUT("/notelp/:id", noTelpHandler.HandleEditNoTelpByID)

	// // Rekening Routes
	// server.GET("/rekening/nasabah/:id", rekeningHandler.HandleGetRekeningByNasabahID)
	// server.GET("/rekening", rekeningHandler.HandleGetRekening)
	// server.GET("/rekening/:id", rekeningHandler.HandleGetRekeningByID)
	// server.POST("/rekening", rekeningHandler.HandleInsertRekening)
	// server.DELETE("/rekening/:id", rekeningHandler.HandleDeleteRekeningByID)
	// server.PUT("/rekening/:id", rekeningHandler.HandleEditRekeningByID)

	// Running in localhost:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
