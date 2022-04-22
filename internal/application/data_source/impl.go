package data_source_application

import (
	common_domain "can-i-eat/internal/domain/common"
	consumer_tag_domain "can-i-eat/internal/domain/consumer_tag"
	food_domain "can-i-eat/internal/domain/food"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

var Impl DataSourceApplication = &dataSourceApplicationImpl{}

type dataSourceApplicationImpl struct {
}

func (d dataSourceApplicationImpl) UploadBhJson(bhList common_domain.BhList) error {
	if len(bhList) == 0 {
		return nil
	}

	foodList := make([]*food_domain.Food, 0)
	foodToTagMap := make(map[string][]*food_tag_domain.FoodTag, 0)
	consumerTagList := make([]*consumer_tag_domain.ConsumerTag, 0)
	for i := range bhList {
		foodID := strings.Join(pinyin.LazyConvert(bhList[i].Name, nil), "_")
		foodList = append(foodList, &food_domain.Food{
			ID:   foodID,
			Name: bhList[i].Name,
		})
		for i2 := range bhList[i].TagList {
			if strings.Contains(bhList[i].TagList[i2], "_") {
				consumerTagName := strings.Split(bhList[i].TagList[i2], "_")[0]
				consumerTagList = append(consumerTagList, &consumer_tag_domain.ConsumerTag{
					ID:   strings.Join(pinyin.LazyConvert(consumerTagName, nil), "_"),
					Name: consumerTagName,
				})
			} else if strings.Contains(bhList[i].TagList[i2], "类") || strings.Contains(bhList[i].TagList[i2], "饮品") {
				foodTagID := strings.Join(pinyin.LazyConvert(bhList[i].TagList[i2], nil), "_")
				foodToTagMap[foodID] = append(foodToTagMap[foodID], &food_tag_domain.FoodTag{
					ID:       foodTagID,
					Name:     bhList[i].TagList[i2],
					ParentID: "fenlei",
				})
			} else {
				foodTagID := strings.Join(pinyin.LazyConvert(bhList[i].TagList[i2], nil), "_")
				foodToTagMap[foodID] = append(foodToTagMap[foodID], &food_tag_domain.FoodTag{
					ID:   foodTagID,
					Name: bhList[i].TagList[i2],
				})
			}

		}
	}
	return nil
}
