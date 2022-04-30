package data_source_facade

import (
	common_domain "can-i-eat/internal/domain/common"
	"can-i-eat/internal/service/data_source"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func handlerUploadBh(c echo.Context) error {
	var bhList common_domain.BhList
	if err := c.Bind(&bhList); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadGateway, err)
	}

	err := data_source_application.Impl.UploadBhJson(bhList)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("导入 %d 条数据", len(bhList)))
}
