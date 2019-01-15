package main

import (
	"chitchat/controllers"
	"chitchat/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(f)
	defer f.Close()

	r := gin.Default()

	//静态文件
	r.Static("/static/", "./public")

	//路由
	authC := controllers.NewAuthController()
	routers.CreateAuthRouters(r, authC)

	threadC := controllers.NewThreadController()
	routers.CreateThreadRouters(r, threadC)

	commonC := controllers.NewCommonController()
	routers.CreateCommonRouters(r, commonC)

	r.Run(":8080")
}
