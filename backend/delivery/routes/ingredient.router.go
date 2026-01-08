package Routes

import (
	"github.com/gin-gonic/gin"

	"soap_factory/delivery/http/handlers"
)

func IngredientRoutes(router *gin.Engine, ingredientHandler *handlers.IngredientHandler) {
	router.GET("/", ingredientHandler.GetAll)
	router.POST("/", ingredientHandler.Create)
	router.GET("/:id", ingredientHandler.GetByID)
	router.PUT("/:id", ingredientHandler.Update)
	router.DELETE("/:id", ingredientHandler.Delete)
}