package food_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	foodGroup := r.Group("/v1/food")
	{
		foodGroup.GET("/list", handlerList)
		foodGroup.POST("/listByFoodTagList", handlerListByFoodTagList)
		foodGroup.POST("/listByFoodTagListAndConsumerTagId", handlerListByFoodTagListAndConsumerTagId)
		foodGroup.GET("/detail", handlerDetail)
		foodGroup.POST("/create", handlerCreate)
		foodGroup.POST("/update", handlerUpdate)
		foodGroup.POST("/delete", handlerDelete)
	}
}
