package data_source_facade

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	dataSourceGroup := r.Group("/v1/dataSource")
	{
		dataSourceGroup.POST("/upload/bh", handlerUploadBh)
	}
}
