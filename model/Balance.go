package model

type Balance struct {
	//model内已经包含id，创建、修改、删除时间, 并且会在对应的操作时自动插入/更新时间
	BaseModel
	UserId  uint `json:"userId"`
	Balance uint `json:"balance"`
}
