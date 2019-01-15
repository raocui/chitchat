package controllers

import (
	"chitchat/utils/session"

	"chitchat/utils"

	"github.com/gin-gonic/gin"
)

func NewCommonController() *CommonController {
	return &CommonController{}
}

type CommonController struct {
}

func (cc *CommonController) Error(c *gin.Context) {
	vals := c.Request.URL.Query()
	check, _ := session.Session(c)
	if check == false {
		utils.GenerateHTML(c.Writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(c.Writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
