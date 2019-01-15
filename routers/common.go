package routers

import (
	"chitchat/controllers"

	"github.com/gin-gonic/gin"
)

func CreateCommonRouters(r *gin.Engine, cc *controllers.CommonController) {
	r.GET("/err", cc.Error)
}
