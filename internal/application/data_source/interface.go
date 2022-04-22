package data_source_application

import (
	common_domain "can-i-eat/internal/domain/common"
)

type DataSourceApplication interface {
	UploadBhJson(list common_domain.BhList) error
}
