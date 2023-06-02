package modeler

import (
	"errors"
	"github.com/jameskeane/bcrypt"
	"github.com/redhoe/couress/utils/simple"
	"gorm.io/gorm"
	"time"
)

type Administrator struct {
	MysqlModel
	Uuid          string     `json:"uuid" gorm:"type:varchar(100);index;comment:uuid;unique"`
	NickName      string     `json:"nick_name" gorm:"column:nick_name;type:char(20);common:昵称"`
	UserName      string     `json:"user_name" gorm:"column:user_name;type:varchar(50);unique"`
	Password      string     `json:"-" gorm:"column:password;type:varchar(100)"`
	Salt          string     `json:"salt" gorm:"column:salt;varchar(50);"`
	RoleId        uint       `json:"role_id" gorm:"column:role_id"`
	Avatar        string     `json:"avatar" gorm:"type:varchar(100);comment:头像地址"`
	Lock          *bool      `json:"lock" gorm:"column:lock;default:false"`
	Token         *string    `json:"-" gorm:";column:token;type:text"`
	LastLoginIp   string     `json:"last_login_ip" gorm:"column:last_login_ip;type:varchar(20)"`
	LastLoginTime *time.Time `json:"last_login_time" gorm:"column:last_login_time"`
	GoogleKey     *string    `json:"-" gorm:"column:google_key;type:varchar(100)"`
}

func (Administrator) TableName() string {
	return "administrator"
}

func NewAdministrator() *Administrator {
	return &Administrator{}
}

func (Administrator) Comment() string {
	return "管理员"
}

func (r *Administrator) GetById(db *gorm.DB, id uint) error {
	err := db.Find(r, "id", id).Error
	if err != nil {
		return err
	}
	if r.Id == 0 {
		return errors.New("user is not exits")
	}
	return nil
}

func (r *Administrator) CheckUserIsValidWithId(db *gorm.DB, id interface{}) error {
	if err := db.Where("lock", false).
		Where("id", id).
		Find(&r).Error; err != nil {
		return err
	}
	if r.Id == 0 {
		return errors.New("user is not exits")
	}
	return nil
}

func (r *Administrator) CheckUserIsValid(db *gorm.DB, userName string) error {
	if err := db.Where("lock", false).
		Where("user_name", userName).
		Find(&r).Error; err != nil {
		return err
	}
	if r.Id == 0 {
		return errors.New("user is not exits")
	}
	return nil
}

func (r *Administrator) CheckPassWord(checkPassWord string) bool {
	checkPassWordEn, _ := bcrypt.Hash(checkPassWord, r.Salt)
	if checkPassWordEn == r.Password {
		return true
	}
	return false
}

func (r *Administrator) EncodePassword(passWord string) (enPassWord, salt string) {
	salt, err := bcrypt.Salt(bcrypt.DefaultRounds)
	if err != nil {
		panic(err)
	}
	enPassWord, err = bcrypt.Hash(passWord, salt)
	if err != nil {
		panic(err)
	}
	return
}

func (r *Administrator) SaveLoginInfo(db *gorm.DB) error {
	return db.Model(&Administrator{}).Where("id", r.Id).
		Updates(map[string]any{
			"last_login_ip":   r.LastLoginIp,
			"last_login_time": r.LastLoginTime,
			"token":           r.Token,
		}).Error
}

func (r *Administrator) DataInit(db *gorm.DB) error {
	// 初始化用户等级
	password, salt := r.EncodePassword("123456")
	reqs := []Administrator{
		{
			NickName:  "超级管理员",
			UserName:  "root",
			Password:  password,
			Salt:      salt,
			GoogleKey: nil,
			Avatar:    "",
			RoleId:    1,
			Uuid:      simple.NewUuid(),
		},
	}
	for _, req := range reqs {
		find := NewAdministrator()
		if stat := db.Model(&Administrator{}).Where("user_name = ?", req.UserName).
			Find(&find).Statement; stat.RowsAffected == 0 {
			if err := db.Create(&req).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

type AdministratorLog struct {
	MysqlModel
	AdminId   int64   `json:"admin_id" gorm:"admin_id"`
	RequestId string  `json:"request_id" gorm:"type:varchar(100)"`
	Message   string  `json:"message" gorm:"type:text"`
	Table     *string `json:"table" gorm:"type:varchar(50)"`
	Action    string  `json:"action" gorm:"type:varchar(20)"`
	Ip        string  `json:"ip" gorm:"type:varchar(20)"`
	UserAgent string  `json:"user_agent" gorm:"type:text"`
	Extends   *string `json:"extends" gorm:"type:json"`
}

func (AdministratorLog) TableName() string {
	return "administrator_log"
}

func NewAdministratorLog() *AdministratorLog {
	return &AdministratorLog{}
}

func (AdministratorLog) Comment() string {
	return "管理员日志"
}

type Role struct {
	MysqlModel
	Name  string `json:"name" gorm:"column:name;type:varchar(100);comment:角色名称"`
	Auths string `json:"auths" gorm:"column:auths;type:text;comment:权限"`
	Desc  string `json:"desc" gorm:"column:desc;type:varchar(200);comment:角色权限描述;"`
}

func (Role) TableName() string {
	return "administrator_role"
}

func NewRole() *Role {
	return &Role{}
}

func (Role) Comment() string {
	return "角色表"
}

func (r *Role) GetById(db *gorm.DB, id interface{}) error {
	err := db.Find(r, "id", id).Error
	if err != nil {
		return err
	}
	if r.Id == 0 {
		return errors.New("IdNotExits")
	}
	return nil
}

func (a *Role) GetAllPage(db *gorm.DB, paging *Paging) ([]Role, error) {
	results := make([]Role, 0)
	var err error
	d := db.Model(&Role{}).Count(&paging.Total)
	paging.GetPages()
	if paging.Total < 1 {
		return results, d.Error
	}
	err = db.Model(&Role{}).Order("created_at desc").Limit(paging.PageSize).Offset(paging.StartNums).Find(&results).Error
	return results, err
}

func (*Role) DataInit(db *gorm.DB) error {
	// 初始化用户等级
	reqs := []Role{
		{
			Name:  "超级管理员",
			Desc:  "管理所有权限",
			Auths: "*",
		},
	}
	for _, req := range reqs {
		find := Role{}
		if stat := db.Model(&Role{}).Where("name", req.Name).
			Find(&find).Statement; stat.RowsAffected == 0 {
			if err := db.Create(&req).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

type Permission struct {
	MysqlModel
	ParentId uint   `json:"parent_id" gorm:"column:parent_id;comment:上级菜单Id"`
	Name     string `json:"name" gorm:"column:name;type:varchar(50);comment:菜单名;unique"`
	Auth     string `json:"auth" gorm:"column:auth;type:varchar(80);comment:权限;unique"`
	Value    string `json:"value" gorm:"column:value;type:varchar(50);comment:扩展名"`
}

func (Permission) TableName() string {
	return "administrator_permission"
}

func NewPermission() *Permission {
	return &Permission{}
}

func (Permission) Comment() string {
	return "权限表"
}

// FindOrCreateAuth 查找创建权限菜单
func (m *Permission) FindOrCreateAuth(db *gorm.DB, auth string) error {
	if err := db.Find(&m, "auth", auth).Error; err != nil {
		return err
	}
	if m.Id == 0 {
		// 创建
		m.ParentId = 1
		m.Auth = auth
		m.Name = auth
		if err := db.Create(&m).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m *Permission) CheckExitsByName(db *gorm.DB, name string) (error, bool) {
	err := db.Find(&m, "name", name).Error
	if err != nil {
		return err, false
	}
	if m.Id == 0 {
		return errors.New("IsNotExits"), false
	}
	return nil, true
}

func (*Permission) DataInit(db *gorm.DB) error {
	// 初始化用户等级
	reqs := []Permission{
		{
			ParentId: 0,
			Name:     "待定义接口",
			Auth:     "undefined apis",
		},
	}
	for _, req := range reqs {
		find := Permission{}
		if stat := db.Model(&Permission{}).Where("name = ?", req.Name).Find(&find).Statement; stat.RowsAffected == 0 {
			if err := db.Create(&req).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

type NodePermission struct {
	Id       uint              `json:"id"`
	ParentId uint              `json:"parent_id"`
	Name     string            `json:"name"`
	Auth     string            `json:"auth"`
	Children []*NodePermission `json:"children"`
}

// MenuTree 递归树[满足无限层级递归]
func (m *Permission) MenuTree(db *gorm.DB) ([]NodePermission, error) {
	menusNodeTree := make([]NodePermission, 0) // 根节点
	// 所有节点
	allMenus := make([]Permission, 0)
	if err := db.Model(&Permission{}).Where("1=1").Find(&allMenus).Error; err != nil {
		return menusNodeTree, err
	}
	for _, m := range allMenus {
		childMenus := make([]*NodePermission, 0)
		if m.ParentId == 0 {
			rootNode := NodePermission{
				Id:       m.Id,
				ParentId: m.ParentId,
				Name:     m.Name,
				Auth:     m.Auth,
				Children: childMenus,
			}
			menusNodeTree = append(menusNodeTree, rootNode)
		}
	}
	for i, _ := range menusNodeTree {
		m.walk(allMenus, &menusNodeTree[i])
	}
	return menusNodeTree, nil
}

// 递归组装菜单树（根节点）
func (m *Permission) walk(allMenus []Permission, rootNode *NodePermission) {
	// 列出所有下级子目录
	nodes := m.childrenList(allMenus, rootNode.Id)
	if len(nodes) == 0 {
		return
	}
	// 遍历这些文件
	for _, node := range nodes {
		newNode := NodePermission{
			Id:       node.Id,
			ParentId: node.ParentId,
			Name:     node.Name,
			Auth:     node.Auth,
			Children: nil,
		}
		m.walk(allMenus, &newNode)
		rootNode.Children = append(rootNode.Children, &newNode)
	}
	return
}

// 获得子节点列表
func (m *Permission) childrenList(allMenus []Permission, pId uint) (menusNodeTree []NodePermission) {
	for _, m := range allMenus {
		if m.ParentId == pId {
			rootNode := NodePermission{
				Id:       m.Id,
				ParentId: m.ParentId,
				Name:     m.Name,
				Auth:     m.Auth,
				Children: nil,
			}
			menusNodeTree = append(menusNodeTree, rootNode)
		}
	}
	return
}
