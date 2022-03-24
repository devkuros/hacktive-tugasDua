package routers

import (
	"tugasdua/configs"
	"tugasdua/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	routers := gin.Default()
	db := configs.ConnectionDB()

	ctrlOrder := controllers.GetOrderDB(db)
	// ctrlItem := controllers.GetItemDB(db)

	// ROUTER ORDER
	routers.POST("/orders", ctrlOrder.CreateDataOrder)
	routers.GET("/orders", ctrlOrder.GetDataOrder)
	routers.DELETE("/orders/:orderId", ctrlOrder.DeleteOrder)
	routers.PUT("/orders/:orderId", ctrlOrder.UpdateDataOrder)

	//ROUTER ITEM
	// routers.POST("/additem", ctrlItem.CreateDataItem)

	return routers
}
