package models

type Nav struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Link      string `json:"link,omitempty"`
	Position  int    `json:"position,omitempty"`
	IsOpennew int    `json:"isOpennew,omitempty"`
	Relation  string `json:"relation,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	Status    int    `json:"status,omitempty"`
	AddTime   int    `json:"addTime,omitempty"`
}

func (Nav) TableName() string {
	return "nav"
}
