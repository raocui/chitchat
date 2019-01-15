package session

import (
	"chitchat/models"

	"fmt"

	"net/http"

	"encoding/base64"

	"time"

	"github.com/gin-gonic/gin"
)

//检查用户是否登入，并且有一个session
func Session(c *gin.Context) (b bool, sess models.Session) {
	cookie, err := c.Cookie("_cookie")

	if err == nil {
		sess = models.Session{Uuid: cookie}
		if ok, _ := (&sess).Check(); !ok {
			b = false
		} else {
			b = true
		}
	}
	fmt.Println("aaaaaaabbbbbbbbb", sess, err)
	return
}

func setFlashMessage(message string, w http.ResponseWriter, r *http.Request) {
	msg := []byte(message)
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}
func showFlashMessage(w http.ResponseWriter, r *http.Request) (msg string, err error) {
	c, err := r.Cookie("flash")
	if err != nil {
		return "", err
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		msg, _ := base64.URLEncoding.DecodeString(c.Value)
		return string(msg), nil
	}

}
