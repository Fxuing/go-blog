package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewBlog(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/new.html", nil)

}
