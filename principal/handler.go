package principal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTrucosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, GetAllTrucos())
}

func CreateTrucoHandler(ctx *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	truco := CreateTruco(input.Name)
	// Actualizar el servidor de replica
	ctx.JSON(http.StatusCreated, truco)
}

func UpdateTrucoHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err = UpdateTruco(id, input.Name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Truco no encontrado"})
		return
	}

	// Actualizar el servidor de replica
	ctx.JSON(http.StatusOK, gin.H{"message": "Truco actualizado"})
}

func DeleteTrucoHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = DeleteTruco(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Truco no encontrado"})
		return
	}

	// Actualizar el servidor de replica 
	ctx.JSON(http.StatusOK, gin.H{"message": "Truco eliminado"})

}