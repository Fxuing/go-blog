package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

/**
Model ===========================================
*/
// 用户
type UserInfo struct {
	gorm.Model
	Username  string `gorm:"size:16"` // 用户名
	Gender    int    // 性别 0 女 1 男
	PhoneCode string `gorm:"size:11"` // 手机号
	PassWord  string `gorm:"size:16"` // 密码
}

// 博客信息
type BlogInfo struct {
	gorm.Model
	Title    string `gorm:"size:255"`
	BlogBody string `gorm:"type:text"`
	Comments []Comments
}

// 评论信息
type Comments struct {
	gorm.Model
	BlogInfoId   int
	UserInfoId   uint
	CommentsBody string   `gorm:"type:text"`
	UserInfo     UserInfo `gorm:"foreignkey:UserInfoId"`
}

// 渠道信息 测试功能
type ChannelInfo struct {
	Id          int
	ChannelName string
	ChannelNo   string
	ParentId    int
}

// 奖品信息
type ActivityAward struct {
	Id          int
	AwardName   string
	AwardTotal  int
	Probability int
	AwardDesc   string
}

/**
初始化数据库，并建立连接
*/
var DB *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	//db, err = gorm.Open("mysql", "root:fxuing#@/go_blog?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", "root:Zopen2013@(120.78.66.185:13306)/go_blog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	/*var tables []interface{}
	tables = append(tables, UserInfo{}, BlogInfo{},Comments{})
	for i := 0; i < len(tables); i++ {
		exist := db.HasTable(tables[i])
		if !exist {
			db.CreateTable(tables[i])
		}
	}*/
	db.AutoMigrate(&UserInfo{}, &BlogInfo{}, &Comments{}, &ChannelInfo{}, &ActivityAward{}, &Student{}, &Class{})
	db.LogMode(true)
	DB = db
	return db, err
}

/**
User db  ===================================
*/
func UserAdd(userinfo UserInfo) error {
	return DB.Create(&userinfo).Error
}
func FindByUserName(userinfo UserInfo) (users []*UserInfo, err error) {
	err = DB.Where("username = ? ", userinfo.Username).Find(&users).Error
	return
}
func FindByUserNameAndPassWord(userinfo UserInfo) (*UserInfo, error) {
	var user = UserInfo{}
	err := DB.Where("username = ? and pass_word = ?", userinfo.Username, userinfo.PassWord).First(&user).Error
	return &user, err
}

/**
BlogInfo db =======================================
*/
func BlogInfoAdd(blogContext BlogInfo) error {
	return DB.Create(&blogContext).Error
}
func BlogLoad() ([]*BlogInfo, error) {
	var userList []*BlogInfo
	err := DB.Order("created_at desc").Find(&userList).Error
	return userList, err
}
func BlogDetail(id int) (*BlogInfo, error) {
	var blogInfo BlogInfo
	var comments []Comments

	err := DB.First(&blogInfo, id).Related(&comments).Error
	for i := 0; i < len(comments); i++ {
		var userInfo UserInfo
		err := DB.Where("id = ?", comments[i].UserInfoId).Find(&userInfo).Error
		if err != nil {
			fmt.Print(err)
		}
		comments[i].UserInfo = userInfo
	}
	blogInfo.Comments = comments
	return &blogInfo, err
}

/**
Comments
*/
func CommentAdd(comment Comments) error {
	return DB.Create(&comment).Error
}

// ----------------
func (channel *ChannelInfo) ChannelAdd(channelInfo ChannelInfo) error {
	return DB.Create(&channelInfo).Error
}
func (channel *ChannelInfo) FindByParentId(id int) (channelInfo []*ChannelInfo, err error) {
	err = DB.Where("parent_id = ?", id).Find(&channelInfo).Error
	return channelInfo, err
}

// -------ActivityAward
func SearchAward() (award []*ActivityAward, err error) {
	err = DB.Find(&award).Error
	return
}
func InventoryReduction(id int) error {
	var lock sync.Mutex
	award := new(ActivityAward)
	award.Id = id
	lock.Lock()
	DB.Find(&award)
	if award.AwardTotal <= 0 {
		return errors.New("库存不足")
	}
	count := award.AwardTotal - 1
	lock.Unlock()
	return DB.Model(&award).Updates(map[string]interface{}{"award_total": count}).Error
}

type Student struct {
	gorm.Model
	Name      string `gorm:"size:64"`
	Age       int    `gorm:"size:3"`
	ClassInfo Class
	ClassId   uint
}

type Class struct {
	gorm.Model
	ClassName string `gorm:"size:64"`
	Students  []Student
}

func SearchClassById(id uint) (*Class, error) {
	class := new(Class)
	class.ID = id
	err := DB.Find(&class).Related(&class.Students).Error
	return class, err
}

func SearchStudentById(id uint) (*Student, error) {
	student := new(Student)
	student.ID = id
	err := DB.Find(&student).Related(&student.ClassInfo).Error
	return student, err
}
