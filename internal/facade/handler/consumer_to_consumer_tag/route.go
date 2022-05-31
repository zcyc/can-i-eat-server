package consumer_to_consumer_tag_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	consumerToConsumerTagGroup := r.Group("/v1/consumerToConsumerTag")
	{
		consumerToConsumerTagGroup.GET("/list", handlerList)
		consumerToConsumerTagGroup.GET("/detail", handlerDetail)
		consumerToConsumerTagGroup.POST("/create", handlerCreate)
		consumerToConsumerTagGroup.POST("/update", handlerUpdate)
		consumerToConsumerTagGroup.POST("/delete", handlerDelete)
	}
}
