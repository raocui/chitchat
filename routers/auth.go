package routers

import (
	"chitchat/controllers"

	"github.com/gin-gonic/gin"
)

func CreateAuthRouters(r *gin.Engine, ac *controllers.AuthController) {
	r.GET("/", ac.Index)
	r.GET("/login", ac.Login)
	r.GET("/logout", ac.Logout)
	r.GET("/signup", ac.SignUp)
	r.POST("/signup_account", ac.SignupAccount)
	r.POST("/authenticate", ac.Authenticate)

}
