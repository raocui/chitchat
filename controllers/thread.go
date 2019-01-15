package controllers

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"chitchat/models"
	"chitchat/utils"
	"chitchat/utils/data"
	"chitchat/utils/session"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ThreadController struct {
}

func NewThreadController() *ThreadController {
	return &ThreadController{}
}

func (tc *ThreadController) NewThread(c *gin.Context) {
	check, _ := session.Session(c)
	if check == false {
		c.Redirect(http.StatusFound, "/login")
	} else {
		tmpl := []string{
			"templates/private.navbar.html",
			"templates/layout.html",
			"templates/new.thread.html",
		}
		var templates *template.Template
		templates = template.Must(template.ParseFiles(tmpl...))
		templates.ExecuteTemplate(c.Writer, "layout", nil)

	}

}

func (tc *ThreadController) CreateThread(c *gin.Context) {

	check, sess := session.Session(c)
	if check == false {
		c.Redirect(http.StatusFound, "/login")
	}

	topic := c.PostForm("topic")
	var thread *models.Thread
	thread = &models.Thread{
		Uuid:      data.CreateUUID(),
		Topic:     topic,
		UserId:    sess.UserId,
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	result := thread.CreateThread()
	if result {
		c.Redirect(http.StatusFound, "/")
	}

}

// /thread/read/:uuid
// 读取单篇帖子 及其回复
func (tc *ThreadController) ReadThread(c *gin.Context) {
	type ThreadUuid struct {
		Uuid string `uri:"uuid" binding:"required"`
	}
	var threadUuid ThreadUuid
	err := c.ShouldBindUri(&threadUuid)
	if err != nil {
		log.Println(err)
	}
	log.Println(threadUuid.Uuid)

	thread, err := models.GetThreadByUuid(threadUuid.Uuid)
	check, sess := session.Session(c)
	log.Println(sess)
	log.Println("thread", thread)
	log.Println("thread id", thread.Id)
	posts := []models.Post{}

	posts, err = thread.Posts()
	log.Println("aaaaaa.posts：", posts)
	var templates *template.Template
	public_tmpl_files := []string{
		"templates/public.thread.html",
		"templates/public.navbar.html",
		"templates/layout.html",
	}
	private_tmpl_files := []string{
		"templates/private.thread.html",
		"templates/private.navbar.html",
		"templates/layout.html",
	}
	if check == false { //未登录
		templates = template.Must(template.ParseFiles(public_tmpl_files...))
	} else { //已登录
		templates = template.Must(template.ParseFiles(private_tmpl_files...))
	}

	templates.ExecuteTemplate(c.Writer, "layout", thread)
}

// /thread/post
//回复帖子
func (tc *ThreadController) PostThread(c *gin.Context) {
	check, sess := session.Session(c)
	if check == false {
		c.Redirect(http.StatusFound, "/login")
	} else {
		uuid := c.PostForm("uuid")
		replyMsg := c.PostForm("body")

		thread, err := models.GetThreadByUuid(uuid)
		if err != nil {
			utils.ErrorMessage(c, "Cannot read thread")
		}
		user, err := sess.User()
		if err != nil {
			log.Println(err)
		}

		_, err = user.CreatePost(thread, replyMsg)
		if err != nil {
			utils.ErrorMessage(c, err.Error())
		}
		c.Redirect(http.StatusFound, "/thread/read/"+thread.Uuid)
	}

}

//创建评论
func (tc *ThreadController) CreateComment(c *gin.Context) {
	check, sess := session.Session(c)
	if check == false {
		c.Redirect(http.StatusFound, "/login")
		os.Exit(1)
	}
	user, err := sess.User()
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/login")
		os.Exit(1)
	}
	comment := models.Comment{}
	content := c.PostForm("content")
	postIdStr := c.PostForm("postId")

	if content == "" {
		log.Println("CreateComment content empty", content)
	}
	postId, err := strconv.Atoi(postIdStr)
	log.Println("aaaaaaaaaaaaaaaaaaaaaaaaaa", content)
	log.Println("bbbbbbbbbbbbb", postId)
	log.Println(err)
	if postId <= 0 {
		log.Println("postId invalid", postId)
	}
	post := models.Post{Id: postId}
	comment = models.Comment{
		Content: content,
		UserId:  user.Id,
		Post:    &post,
	}
	err = comment.Create()
	if err != nil {
		log.Println(err)
	}

}

func (tc *ThreadController) CommentList(c *gin.Context) {

	postIdStr := c.Query("id")
	postId, _ := strconv.Atoi(postIdStr)
	posts, _ := models.GetPost(postId)
	log.Println("ccccccc", postId)
	log.Println(posts)
	var templates *template.Template
	public_tmpl_files := []string{
		"templates/private.comment.html",
		"templates/private.navbar.html",
		"templates/layout.html",
	}

	templates = template.Must(template.ParseFiles(public_tmpl_files...))

	templates.ExecuteTemplate(c.Writer, "layout", posts)
}
