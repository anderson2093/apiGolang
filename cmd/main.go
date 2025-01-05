package main

import (
	_ "github.com/anderson2093/apiGolang/cmd/docs"
	"github.com/anderson2093/apiGolang/config"
	"github.com/anderson2093/apiGolang/interfaces/controllers"
	"github.com/anderson2093/apiGolang/interfaces/repositories"
	"github.com/anderson2093/apiGolang/interfaces/routes"
	"github.com/anderson2093/apiGolang/usecase"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Categories API
// @version 1.0
// @description API para gestionar categorías en una tienda virtual
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
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

	// Ruta Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Registra las rutas de la API
	routes.RegisterRoutes(router, controller)

	// Iniciar servidor
	router.Run(":8080")
}
