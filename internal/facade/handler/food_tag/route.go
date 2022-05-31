package food_tag_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	foodTagGroup := r.Group("/v1/foodTag")
	{
		foodTagGroup.GET("/list", handlerList)
		foodTagGroup.GET("/detail", handlerDetail)
		foodTagGroup.POST("/create", handlerCreate)
		foodTagGroup.POST("/update", handlerUpdate)
		foodTagGroup.POST("/delete", handlerDelete)
	}
}
