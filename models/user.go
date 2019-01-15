package models

import (
	"chitchat/utils/data"
	"log"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	UpdatedAt int64
	CreatedAt int64
}

func UserByEmail(email string) (user User, err error) {
	err = Db.QueryRow("select id, uuid,email,password,created_at from users where email=?", email).Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Println(err.Error())
	}
	return

}

func (user *User) CreateUser() (bool, error) {
	result, err := Db.Exec("INSERT INTO users (uuid,`name`,email,`password`,updated_at,created_at) VALUES (?,?,?,?,?,?)", &user.Uuid, &user.Name, &user.Email, data.Encrypt(user.Password), &user.UpdatedAt, &user.CreatedAt)

	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Println(result)
	return true, nil
}
