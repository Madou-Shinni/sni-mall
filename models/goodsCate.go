package models

type GoodsCate struct {
	Id             int         `json:"id,omitempty"`
	Title          string      `json:"title,omitempty"`
	CateImg        string      `json:"cateImg,omitempty"`
	Link           string      `json:"link,omitempty"`
	Template       string      `json:"template,omitempty"`
	Pid            int         `json:"pid,omitempty"`
	SubTitle       string      `json:"subTitle,omitempty"`
	Keywords       string      `json:"keywords,omitempty"`
	Description    string      `json:"description,omitempty"`
	Sort           int         `json:"sort,omitempty"`
	Status         int         `json:"status,omitempty"`
	AddTime        int         `json:"addTime,omitempty"`
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid;references:Id" json:"goodsCateItems,omitempty"`
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
