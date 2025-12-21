package handlers

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    
    "soap_factory/application/services"
)

type IngredientHandler struct {
    service *services.IngredientService
}

func NewIngredientHandler(service *services.IngredientService) *IngredientHandler {
    return &IngredientHandler{service: service}
}

// Handler methods 
func (h *IngredientHandler) Create(c *gin.Context) {
    // TODO: Implement
    c.JSON(http.StatusOK, gin.H{"message": "Create ingredient"})
}

func (h *IngredientHandler) GetAll(c *gin.Context) {
    // TODO: Implement
    c.JSON(http.StatusOK, gin.H{"message": "Get all ingredients"})
}

func (h *IngredientHandler) GetByID(c *gin.Context) {
    // TODO: Implement
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "Get ingredient", "id": id})
}

func (h *IngredientHandler) Update(c *gin.Context) {
    // TODO: Implement
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "Update ingredient", "id": id})
}

func (h *IngredientHandler) Delete(c *gin.Context) {
    // TODO: Implement
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "Delete ingredient", "id": id})
}

func (h *IngredientHandler) RecordPurchase(c *gin.Context) {
    // TODO: Implement
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"message": "Record purchase", "ingredient_id": id})
}

func (h *IngredientHandler) GetLowStock(c *gin.Context) {
    // TODO: Implement
    c.JSON(http.StatusOK, gin.H{"message": "Get low stock ingredients"})
}