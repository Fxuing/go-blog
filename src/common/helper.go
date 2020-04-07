package common

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ContextUserKey = "UserInfo"
	Session        = "gin-session"
)

func SetSession(key string, value interface{}, c *gin.Context) {
	session := sessions.Default(c)
	session.Set(key, value)
	err := session.Save()
	if err != nil {
		fmt.Println("set session ", err)
	}
}

func GetSession(key string, c *gin.Context) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

func DelSession(key string, c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(key)
	err := session.Save()
	if err != nil {
		fmt.Println("delete session ", err)
	}
}
