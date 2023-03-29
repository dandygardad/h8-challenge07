package main

import (
	"challenge07/config"
	"challenge07/repository"
	"challenge07/routes"
	"challenge07/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	err := config.InitPostgres()
	if err != nil {
		panic(err)
	}

	newRepo := repository.NewRepo(config.PSQL.DB)
	newService := service.NewService(newRepo)
	router := gin.New()
	routes.BookRouter(router, newService)

	port := os.Getenv("PORT")
	err = router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
