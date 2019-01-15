package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Falied to open log file", err)
	}
	Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

//for logging
//func Info(args ...interface{}) {
//	logger.SetPrefix("INFO ")
//	logger.Println(args...)
//}
//
//func Danger(args ...interface{}) {
//	logger.SetPrefix("ERROR ")
//	logger.Println(args)
//}
//func Warning(args ...interface{}) {
//	logger.SetPrefix("WARNING ")
//	logger.Println(args...)
//}
