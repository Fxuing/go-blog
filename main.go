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
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var db, err = model.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	router := gin.Default()
	//router.LoadHTMLGlob(GetCurrentPath()+"/src/go-blog/templates/**/*")
	//router.Static("/static", GetCurrentPath()+"/src/go-blog/static")
	router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "/src/go-blog/templates/**/*"))
	router.Static("/static", filepath.Join(os.Getenv("GOPATH"), "/src/go-blog/static"))

	gob.Register(model.UserInfo{})
	store := cookie.NewStore([]byte(common.ContextUserKey))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 3600 * 24, Path: "/"})
	router.Use(sessions.Sessions(common.Session, store))

	router.GET("/", controller.Index)

	router.GET("/login", controller.LoginGet)
	router.POST("/login", controller.LoginPost)
	router.GET("/logout", controller.Logout)

	router.GET("/register", controller.Register)
	router.POST("/register", controller.RegisterPost)

	router.GET("/blog/new", controller.NewBlog)

	router.Run(":8000")
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Print(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
