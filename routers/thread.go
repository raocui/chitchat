package routers

import (
	"chitchat/controllers"

	"github.com/gin-gonic/gin"
)

func CreateThreadRouters(r *gin.Engine, tc *controllers.ThreadController) {
	r.GET("/thread/new", tc.NewThread)
	r.POST("/thread/create", tc.CreateThread)

	r.GET("/thread/read/:uuid", tc.ReadThread)
	r.POST("/thread/post", tc.PostThread)

	r.POST("/post/comment", tc.CreateComment)
	r.GET("/post/comment/list", tc.CommentList)
}
