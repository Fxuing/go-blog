package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/src/common"
	"go-blog/src/model"
	"net/http"
)

func Index(c *gin.Context) {

	user := common.GetSession(common.ContextUserKey, c)
	userInfo, _ := user.(model.UserInfo)
	fmt.Println("===============userinfo:", userInfo)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"user": userInfo.Username,
	})
}
