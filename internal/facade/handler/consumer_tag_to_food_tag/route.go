package consumer_tag_to_food_tag_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	consumerTagToFoodTagGroup := r.Group("/v1/consumerTagToFoodTag")
	{
		consumerTagToFoodTagGroup.GET("list", handlerList)
		consumerTagToFoodTagGroup.GET("detail", handlerDetail)
		consumerTagToFoodTagGroup.GET("listByConsumerTagAndEatMode", handlerListByConsumerTagAndEatMode)
		consumerTagToFoodTagGroup.POST("create", handlerCreate)
		consumerTagToFoodTagGroup.POST("update", handlerUpdate)
		consumerTagToFoodTagGroup.POST("delete", handlerDelete)
	}
}
