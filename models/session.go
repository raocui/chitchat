package models

import (
	"chitchat/utils/data"
	"log"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	UpdatedAt int64
	CreatedAt int64
}

func (user *User) CreateSession() (session Session, err error) {
	stmt, err := Db.Prepare("INSERT INTO sessions(uuid,email,user_id,updated_at,created_at) Values(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	uuid := data.CreateUUID()
	res, err := stmt.Exec(uuid, user.Email, user.Id, user.UpdatedAt, user.CreatedAt)
	if err != nil {
		log.Println(err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	if err == nil {
		session = Session{
			int(id), uuid, user.Email, user.Id, user.UpdatedAt, user.CreatedAt,
		}
	}

	return

}

func (sess *Session) DeleteByUUID() (err error) {
	stmt, err := Db.Prepare("DELETE FROM sessions WHERE uuid=?")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(sess.Uuid)
	if err != nil {
		log.Println(err)
	}
	return
}

// Check if session is valid in the database
func (sess *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", sess.Uuid).
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if sess.Id != 0 {
		valid = true
	}
	log.Println("check", valid)
	return
}

func (sess *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid,name,email,created_at,updated_at from users where id=?", sess.UserId).Scan(&user.Id, &user.Uuid, &user.Email, &user.Name, &user.UpdatedAt, &user.CreatedAt)
	return
}
