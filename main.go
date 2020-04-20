package main

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-blog/src/common"
	"go-blog/src/controller"
	"go-blog/src/model"
	"os"
	"path/filepath"
)

func main() {
	var db, err = model.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	router := gin.Default()

	router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "/src/go-blog/templates/**/*"))
	router.Static("/static", filepath.Join(os.Getenv("GOPATH"), "/src/go-blog/static"))

	gob.Register(model.UserInfo{})

	setGinSession(router)

	router.GET("/", controller.Index)

	router.GET("/login", controller.LoginGet)
	router.POST("/login", controller.LoginPost)
	router.GET("/logout", controller.Logout)

	router.GET("/register", controller.Register)
	router.POST("/register", controller.RegisterPost)

	router.GET("/blog/new", controller.NewBlog)
	router.POST("/blog/saveBlog", controller.SaveBlog)
	router.GET("/blog/detail/:id", controller.DetailBlog)

	router.POST("/blog/comment", controller.Comment)

	router.Run(":8000")
}

func setGinSession(router *gin.Engine) {
	store := cookie.NewStore([]byte(common.ContextUserKey))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 3600 * 24, Path: "/"})
	router.Use(sessions.Sessions(common.Session, store))
}
