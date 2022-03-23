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
	routers.POST("/addorder", ctrlOrder.CreateDataOrder)

	//ROUTER ITEM
	// routers.POST("/additem", ctrlItem.CreateDataItem)

	return routers
}
