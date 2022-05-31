package eat_mode_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	eatModeGroup := r.Group("/v1/eatMode")
	{
		eatModeGroup.GET("list", handlerList)
		eatModeGroup.GET("detail", handlerDetail)
		eatModeGroup.POST("create", handlerCreate)
		eatModeGroup.POST("update", handlerUpdate)
		eatModeGroup.POST("delete", handlerDelete)
	}
}
