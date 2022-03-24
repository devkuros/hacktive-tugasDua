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

func (idb *orderDB) GetDataOrder(ctx *gin.Context) {
	var getOrder []models.Order

	if err := idb.DB.Find(&getOrder).Error; err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": getOrder,
	})
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

func (idb *orderDB) DeleteOrder(ctx *gin.Context) {
	var deleteOrder models.Order

	getId := ctx.Params.ByName("orderId")

	if err := idb.DB.Where("order_id = ?", getId).Delete(&deleteOrder).Error; err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id #" + getId: "deleted",
	})
}

func (idb *orderDB) UpdateDataOrder(ctx *gin.Context) {
	var updateOrder models.Order

	getId := ctx.Params.ByName("orderId")
	if err := idb.DB.Where("order_id = ?", getId).First(&updateOrder).Error; err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := ctx.ShouldBindJSON(&updateOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	idb.DB.Save(&updateOrder)
	ctx.JSON(http.StatusOK, gin.H{
		"data": updateOrder,
	})

}
