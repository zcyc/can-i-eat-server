package data_source_application

import (
	"can-i-eat/common/constant"
	common_domain "can-i-eat/internal/domain/common"
	consumer_tag_domain "can-i-eat/internal/domain/consumer_tag"
	food_domain "can-i-eat/internal/domain/food"
	food_tag_domain "can-i-eat/internal/domain/food_tag"
	consumer_tag_service "can-i-eat/internal/service/consumer_tag"
	consumer_tag_to_food_tag_service "can-i-eat/internal/service/consumer_tag_to_food_tag"
	food_service "can-i-eat/internal/service/food"
	food_tag_service "can-i-eat/internal/service/food_tag"
	food_to_food_tag_service "can-i-eat/internal/service/food_to_food_tag"
	"github.com/labstack/gommon/log"
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

	// 设置获取拼音的时候保留字母
	a := pinyin.NewArgs()
	a.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}

	// 开始处理数据
	foodList := make([]*food_domain.Food, 0)
	foodToFoodTagMap := make(map[string][]string, 0)
	foodTagList := make([]*food_tag_domain.FoodTag, 0)
	consumerTagList := make([]*consumer_tag_domain.ConsumerTag, 0)
	foodToConsumerTagMap := make(map[string][]string, 0)
	for i := range bhList {
		foodID := strings.Join(pinyin.LazyConvert(bhList[i].Name, &a), "_")
		foodList = append(foodList, &food_domain.Food{
			Active: constant.Activated,
			Flag:   constant.Normal,
			ID:     foodID,
			Name:   bhList[i].Name,
		})
		for i2 := range bhList[i].TagList {
			if strings.Contains(bhList[i].TagList[i2], "_") {
				foodToConsumerTagMap[foodID] = append(foodToConsumerTagMap[foodID], bhList[i].TagList[i2])
				consumerTagName := strings.Split(bhList[i].TagList[i2], "_")[0]
				if isConsumerTagExist(consumerTagList, consumerTagName) == true {
					continue
				}
				consumerTagList = append(consumerTagList, &consumer_tag_domain.ConsumerTag{
					Active: constant.Activated,
					Flag:   constant.Normal,
					ID:     strings.Join(pinyin.LazyConvert(consumerTagName, &a), "_"),
					Name:   consumerTagName,
				})
			} else if strings.Contains(bhList[i].TagList[i2], "类") || strings.Contains(bhList[i].TagList[i2], "饮品") {
				foodTagID := strings.Join(pinyin.LazyConvert(bhList[i].TagList[i2], &a), "_")
				foodTag := &food_tag_domain.FoodTag{
					Active:   constant.Activated,
					Flag:     constant.Normal,
					ID:       foodTagID,
					Name:     bhList[i].TagList[i2],
					ParentID: "fenlei",
				}
				foodToFoodTagMap[foodID] = append(foodToFoodTagMap[foodID], foodTagID)
				if isFoodTagExist(foodTagList, bhList[i].TagList[i2]) == true {
					continue
				}
				foodTagList = append(foodTagList, foodTag)
			} else {
				foodTagID := strings.Join(pinyin.LazyConvert(bhList[i].TagList[i2], &a), "_")
				foodTag := &food_tag_domain.FoodTag{
					Active: constant.Activated,
					Flag:   constant.Normal,
					ID:     foodTagID,
					Name:   bhList[i].TagList[i2],
				}
				foodToFoodTagMap[foodID] = append(foodToFoodTagMap[foodID], foodTagID)
				if isFoodTagExist(foodTagList, bhList[i].TagList[i2]) == true {
					continue
				}
				foodTagList = append(foodTagList, foodTag)
			}
		}
	}

	// 批量导入食物
	err := food_service.Impl.BatchCreate(foodList)
	if err != nil {
		return err
	}

	// 批量导入用户标签
	err = consumer_tag_service.Impl.BatchCreate(consumerTagList)
	if err != nil {
		return err
	}
	log.Info("批量导入用户标签成功")

	// 批量导入食物标签
	err = food_tag_service.Impl.BatchCreate(foodTagList)
	if err != nil {
		return err
	}
	log.Info("批量导入食物标签成功")

	// 批量绑定食物和食物标签
	err = food_to_food_tag_service.Impl.Bind(foodTagList, foodToFoodTagMap)
	if err != nil {
		return err
	}
	log.Info("批量绑定食物和食物标签成功")

	// 批量绑定食物标签和用户标签
	err = consumer_tag_to_food_tag_service.Impl.Bind(foodToFoodTagMap, foodToConsumerTagMap)
	if err != nil {
		return err
	}
	log.Info("批量绑定食物标签和用户标签成功")

	return nil
}

func isConsumerTagExist(consumerTagList []*consumer_tag_domain.ConsumerTag, tagName string) bool {
	for i := range consumerTagList {
		if consumerTagList[i].Name == tagName {
			return true
		}
	}
	return false
}

func isFoodTagExist(consumerTagList []*food_tag_domain.FoodTag, tagName string) bool {
	for i := range consumerTagList {
		if consumerTagList[i].Name == tagName {
			return true
		}
	}
	return false
}
