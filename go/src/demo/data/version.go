package data

type AddVersionReq struct {
	// 名称
	Name string `gorm:"column:name;type:varchar(128);index:version_name,unique" json:"name" column:"name"`
	//描述
	Desc string `gorm:"column:desc;type:varchar(512)" json:"desc" column:"desc"`
	// 是否是发行版本
	IsRelease int `gorm:"column:is_release;type:int" json:"isRelease" column:"is_release"` // 是否发行版本：1，是，2 否
	// 版本计划开始时间
	PlanStartTime string `gorm:"column:plan_start_time;type:datetime" json:"planStartTime"`
	// 版本计划结束时间
	PlanEndTime string `gorm:"column:plan_end_time;type:datetime" json:"planEndTime" `
}

// 小版本新增
type AddSmallVersionReq struct {
	Vid  int64  `json:"vid"` // 版本id
	Name string `json:"Name"`
}

type RenameSmallVersionRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"Name"`
}

type ListVersionRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	// 是否是发行版本
	IsRelease int ` json:"isRelease" ` // 1 发行，2 未发行
	// 版本计划开始时间
	PlanStartTime string `json:"planStartTime"`
	// 版本计划结束时间
	PlanEndTime string `json:"planEndTime"`
}

// 小版本删除
type DeleteSmallVersionRequest struct {
	Id int64 `json:"id"`
}

// 版本编辑

type EditVersionRequest struct {
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	//描述
	Desc string `json:"desc" `
	// 是否是发行版本
	IsRelease int ` json:"isRelease" ` // 1 发行，2 未发行
	// 版本计划开始时间
	PlanStartTime string `json:"planStartTime"`
	// 版本计划结束时间
	PlanEndTime string `json:"planEndTime"`
}

// 版本删除
type DeleteVersionRequest struct {
	Id int64 `json:"id"`
}
