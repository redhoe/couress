package modeler

type IssueTag struct {
	MysqlModel
	MysqlDeleteModel
	Name string `json:"name" gorm:""`
	Show bool   `json:"show" gorm:""`
	Sort int64  `json:"sort" gorm:"default:9999"`
}

func (*IssueTag) TableName() string {
	return "issue_tag"
}

func (*IssueTag) Comment() string {
	return "反馈信息标签"
}

func NewIssueTag() *IssueTag {
	return &IssueTag{}
}

type Issue struct {
	MysqlModel
	MysqlDeleteModel
	IdentityId uint      `json:"identity_id" gorm:""`
	TagId      uint      `json:"tag_id" gorm:""`
	Tag        *IssueTag `json:"tag" gorm:"foreignKey:tag_id;references:id;comment:关联1对1"`
	Content    string    `json:"content" gorm:""`
	Image      string    `json:"image" gorm:"type:text"`
	Contact    string    `json:"contact" gorm:""`
	Status     int       `json:"status" gorm:""`
	Reply      bool      `json:"reply" gorm:"default:false"`
}

func (Issue) TableName() string {
	return "issue"
}

func (Issue) Comment() string {
	return "反馈信息"
}
func NewIssue() *Issue {
	return &Issue{}
}

type IssueMessage struct {
	MysqlModel
	MysqlDeleteModel
	IssueId uint   `json:"issue_id" gorm:""`
	Message string `json:"message" gorm:""`
	AdminId *uint  `json:"admin_id" gorm:""`
}

func (IssueMessage) TableName() string {
	return "issue_message"
}

func (IssueMessage) Comment() string {
	return "反馈回复信息"
}
func NewIssueMessage() *IssueMessage {
	return &IssueMessage{}
}
