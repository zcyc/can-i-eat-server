package food_to_food_tag_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	foodToFoodTagGroup := r.Group("/v1/foodToFoodTag")
	{
		foodToFoodTagGroup.GET("/list", handlerList)
		foodToFoodTagGroup.GET("/detail", handlerDetail)
		foodToFoodTagGroup.POST("/create", handlerCreate)
		foodToFoodTagGroup.POST("/update", handlerUpdate)
		foodToFoodTagGroup.POST("/delete", handlerDelete)
	}
}
