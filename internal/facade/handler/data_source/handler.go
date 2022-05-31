package data_source_facade

import (
	common_domain "can-i-eat/internal/domain/common"
	"can-i-eat/internal/service/data_source"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

func handlerUploadBh(c *gin.Context) {
	var bhList common_domain.BhList
	if err := c.Bind(&bhList); err != nil {
		log.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	err := data_source_application.Impl.UploadBhJson(bhList)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("导入 %d 条数据", len(bhList)))
	return
}
