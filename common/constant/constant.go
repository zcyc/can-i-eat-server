package constant

const (
	/*
		删除状态
		0: 未删除
		1: 已删除
	*/
	DataNormal  = 0
	DataDeleted = 1
	/*
		上下架状态
		0: 下架
		1: 上架
	*/
	DataInactivated = 0
	DataActivated   = 1
)

const (
	/*
		用户标签和食品标签的关系
	*/
	EatModeRecommend = "tui_jian_shi_yong"
	EatModeWarning   = "jin_shen_shi_yong"
)
