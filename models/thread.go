package models

import (
	"chitchat/utils/data"
	"errors"
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	UpdatedAt int64
	CreatedAt int64
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	Comments  []Comment
	UpdatedAt int64
	CreatedAt int64
}

type Comment struct {
	Id      int
	Content string
	UserId  int
	Post    *Post
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id,uuid,topic,user_id,created_at from threads")
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			log.Println(err)
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return

}

func (thread Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id,uuid,body,user_id,thread_id,created_at,updated_at from posts where thread_id=?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.UpdatedAt, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}

	rows.Close()
	return
}

//获得某篇回复
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}

	err = Db.QueryRow("select id,uuid,body,user_id,thread_id from posts where id=?", id).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId)

	rows, err := Db.Query("select id,content,user_id from comments where post_id=?", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.UserId)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}
func (p *Post) CreatedAtDate() string {
	t := time.Unix(p.CreatedAt, 0)
	return t.String()
}

func (p *Post) User() (user User) {
	user = User{}
	Db.QueryRow("select id,uuid, name, email,password,updated_at,created_at from users where id=?", p.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.UpdatedAt, &user.Password)
	return
}
func GetThreadByUuid(uuid string) (thread Thread, err error) {
	thread = Thread{}
	err = Db.QueryRow("select id,uuid,topic,user_id,updated_at,created_at  from threads where uuid=?", uuid).Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.UpdatedAt, &thread.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return
}
func (thread Thread) CreateThread() bool {
	result, err := Db.Exec("INSERT INTO threads (uuid,topic,user_id,updated_at,created_at) VALUES (?,?,?,?,?)", &thread.Uuid, &thread.Topic, &thread.UserId, &thread.UpdatedAt, &thread.CreatedAt)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(result)
	return true
}

func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {

	statement := "INSERT INTO posts (uuid,body,user_id,thread_id,created_at,updated_at) VALUES (?,?,?,?,?,?) "
	stmt, err := Db.Prepare(statement)
	log.Println(stmt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(data.CreateUUID(), body, user.Id, conv.Id, time.Now().Unix(), time.Now().Unix())
	lastId, _ := result.LastInsertId()
	log.Println(lastId)
	stmt2, err := Db.Prepare("select id,uuid,body,user_id,thread_id,updated_at,created_at from posts where id=?")
	err = stmt2.QueryRow(lastId).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.UpdatedAt, &post.CreatedAt)
	log.Println("stmt2:", stmt2)
	defer stmt2.Close()
	log.Println("aaaaaaaaaaaaaaaaaaaaa")
	log.Println("post=", post)
	log.Println(err)
	return
}

func (t Thread) CreatedAtDate() string {

	d := time.Unix((t).CreatedAt, 0)

	return d.String()
}

func (t *Thread) NumReplies() int64 {
	var cnt int64

	row := Db.QueryRow("select count(*) from posts where thread_id=?", t.Id).Scan(&cnt)
	log.Println(row)
	log.Printf("%T \t %v", cnt, cnt)
	return cnt
}

// Get the user who started this thread
func (thread Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

//创建评论
func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}

	rs, err := Db.Exec("insert into comments (content,user_id,post_id) values (?,?,?)", comment.Content, comment.UserId, comment.Post.Id)
	log.Println(err)
	log.Println(rs)
	return
}
