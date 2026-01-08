package http

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"soap_factory/infrastructure/repositories"
	"soap_factory/application/services"
	"soap_factory/delivery/handlers"
)

func Routes(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	// router.Use(corsMiddleware())

	// health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	ingredientRepo := repositories.NewIngredientRepository(db)
	ingredientService := services.NewIngredientService(ingredientRepo)
	ingredientHandler := NewIngredientHandler(ingredientService)

	ingredientRoute := router.Group("/ingredients")
	IngredientRoutes(ingredientRoute, ingredientHandler)
	return router
}