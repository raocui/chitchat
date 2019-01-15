package utils

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Convenience function to redirect to the error message page
func ErrorMessage(c *gin.Context, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(c.Writer, c.Request, strings.Join(url, ""), 302)
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
