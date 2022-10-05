package models

type Access struct {
	Id          int      `json:"id,omitempty"`
	ModuleName  string   `json:"moduleName,omitempty"` //模块名称
	ActionName  string   `json:"actionName,omitempty"` //操作名称
	Type        int      `json:"type,omitempty"`       //节点类型 :  1、表示模块    2、表示菜单     3、操作
	Url         string   `json:"url,omitempty"`        //路由跳转地址
	ModuleId    int      `json:"moduleId,omitempty"`   //此module_id和当前模型的id关联       module_id= 0 表示模块
	Sort        int      `json:"sort,omitempty"`
	Description string   `json:"description,omitempty"`
	Status      int      `json:"status,omitempty"`
	AddTime     int      `json:"addTime,omitempty"`
	AccessItem  []Access `gorm:"foreignKey:ModuleId;references:Id" json:"accessItem,omitempty"`

	Checked bool `gorm:"-" json:"checked,omitempty"` // 忽略本字段
}

func (Access) TableName() string {
	return "access"
}
