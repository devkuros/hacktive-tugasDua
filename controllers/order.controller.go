package controllers

import (
	"net/http"
	"tugasdua/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderDB struct {
	DB *gorm.DB
}

func GetOrderDB(db *gorm.DB) *orderDB {
	return &orderDB{
		DB: db,
	}
}

func (idb *orderDB) CreateDataOrder(ctx *gin.Context) {
	var createOrder models.Order

	if err := ctx.ShouldBindJSON(&createOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{OrderedAt: createOrder.OrderedAt, CostumerName: createOrder.CostumerName, ItemId: createOrder.ItemId}
	idb.DB.Create(&order)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": order,
	})
}
