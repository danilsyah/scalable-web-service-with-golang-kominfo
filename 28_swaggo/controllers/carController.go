package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Car represents the model for an car
type Car struct {
	CarID string `json:"car_id" example:"1"`
	Brand string `json:"brand" example:"Honda"`
	Model string `json:"model" example:"CRV"`
	Price int    `json:"price" example:"50000000"`
}

var CarDatas = []Car{}

// CreateCar godoc
// @summary Create a New Car
// @Description Create a new car with the input car
// @Tags Cars
// @Accept json
// @Produce json
// @Param car body Car true "Create car"
// @Success 200 {object} Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}
