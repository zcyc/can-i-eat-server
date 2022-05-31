package consumer_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	consumerGroup := r.Group("/v1/consumer")
	{
		consumerGroup.GET("/list", handlerList)
		consumerGroup.GET("/detail", handlerDetail)
		consumerGroup.POST("/create", handlerCreate)
		consumerGroup.POST("/update", handlerUpdate)
		consumerGroup.POST("/delete", handlerDelete)
	}
}
