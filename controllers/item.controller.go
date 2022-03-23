package controllers

import (
	"net/http"
	"tugasdua/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type itemDB struct {
	DB *gorm.DB
}

func GetItemDB(db *gorm.DB) *itemDB {
	return &itemDB{
		DB: db,
	}
}

func (i *itemDB) CreateDataItem(ctx *gin.Context) {
	var dataItem models.Item

	if err := ctx.ShouldBindJSON(&dataItem); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	item := models.Item{ItemCode: dataItem.ItemCode, Descryption: dataItem.Descryption, Quantity: dataItem.Quantity}
	i.DB.Create(&item)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": item,
	})
}
