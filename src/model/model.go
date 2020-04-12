package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
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
	BlogInfoId   uint
	UserId       uint
	UserInfo     UserInfo
	CommentsBody string `gorm:"type:text"`
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
	db.AutoMigrate(&UserInfo{}, &BlogInfo{}, &Comments{})
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
	blogInfo.Comments = comments
	return &blogInfo, err
}

/**
Comments
*/
