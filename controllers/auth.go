package controllers

import (
	"chitchat/models"
	"chitchat/utils/data"
	"chitchat/utils/session"
	"html/template"
	"log"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthController() *AuthController {
	return &AuthController{}
}

type AuthController struct {
}

func (authc *AuthController) Index(c *gin.Context) {
	threads, err := models.Threads()
	log.Println("threads list:", threads)

	if err == nil {
		check, _ := session.Session(c)
		public_tmpl_files := []string{
			"templates/index.html",
			"templates/layout.html",
			"templates/public.navbar.html",
		}
		private_tmpl_files := []string{
			"templates/index.html",
			"templates/layout.html",
			"templates/private.navbar.html",
		}
		var templates *template.Template
		if check == true { //登录了
			templates = template.Must(template.ParseFiles(private_tmpl_files...))

		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))

		}

		templates.ExecuteTemplate(c.Writer, "layout", threads)

	}
}

func (authc *AuthController) Login(c *gin.Context) {

	tmpl := []string{
		"templates/login.layout.html",
		"templates/login.html",
	}

	var templates *template.Template
	templates = template.Must(template.ParseFiles(tmpl...))

	templates.ExecuteTemplate(c.Writer, "layout", nil)
}

func (authc *AuthController) Logout(c *gin.Context) {
	cookie, err := c.Cookie("_cookie")
	if err == nil {
		session := &models.Session{Uuid: cookie}
		session.DeleteByUUID()
	}
	c.Redirect(http.StatusFound, "/")

}

func (authc *AuthController) SignUp(c *gin.Context) {
	tmpl := []string{
		"templates/login.layout.html",
		"templates/signup.html",
	}

	var templates *template.Template
	templates = template.Must(template.ParseFiles(tmpl...))
	templates.ExecuteTemplate(c.Writer, "layout", nil)
}

func (authc *AuthController) SignupAccount(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := &models.User{
		Uuid:      data.CreateUUID(),
		Name:      name,
		Email:     email,
		Password:  password,
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	result, _ := user.CreateUser()
	log.Println(result)
	//c.Redirect(http.StatusFound,"/login")
	http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)

}

func (authc *AuthController) Authenticate(c *gin.Context) {
	log.Println(c.PostForm("email"), c.PostForm("password"))
	user, err := models.UserByEmail(c.PostForm("email"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("user:", user)
	user.UpdatedAt = time.Now().Unix()
	user.CreatedAt = time.Now().Unix()
	if user.Password == data.Encrypt(c.PostForm("password")) {
		session, err := (&user).CreateSession()
		if err == nil {
			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.Uuid,
				HttpOnly: true,
			}
			http.SetCookie(c.Writer, &cookie)
			c.Redirect(http.StatusFound, "/")
		} else {
			c.Redirect(http.StatusFound, "/login")
		}

	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
