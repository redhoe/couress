package modeler

type Poster struct {
	MysqlModel
	MysqlDeleteModel
	Image string `json:"image"`
	Show  bool   `json:"show"`
	Lang  string `json:"lang" gorm:"type:varchar(10)"`
	Sort  int    `json:"sort" gorm:"type:int(4);default:9999"` // 排序，由小到大
}

func (*Poster) TableName() string {
	return "poster"
}

func (*Poster) Comment() string {
	return "海报"
}

func NewPoster() *Poster {
	return &Poster{}
}
