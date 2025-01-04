package main

import (
	"github.com/anderson2093/apiGolang/config"
	"github.com/anderson2093/apiGolang/interfaces/controllers"
	"github.com/anderson2093/apiGolang/interfaces/repositories"
	"github.com/anderson2093/apiGolang/interfaces/routes"
	"github.com/anderson2093/apiGolang/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDatabase()
	db := config.DB.Database("categoriesDB")

	// Iniciar dependencias
	repo := repositories.NewMongoCategoryRepository(db)
	usecase := usecase.NewCategoryUsecase(repo)
	controller := controllers.NewCategoryController(usecase)

	// Configurar rutas
	router := gin.Default()
	routes.RegisterRoutes(router, controller)

	// Iniciar servidor
	router.Run(":8080")
}
