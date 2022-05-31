package consumer_tag_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	consumerTagGroup := r.Group("/v1/consumerTag")
	{
		consumerTagGroup.GET("/list", handlerList)
		consumerTagGroup.GET("/detail", handlerDetail)
		consumerTagGroup.POST("/create", handlerCreate)
		consumerTagGroup.POST("/update", handlerUpdate)
		consumerTagGroup.POST("/delete", handlerDelete)
	}
}
