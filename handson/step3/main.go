package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"gopkg.in/go-playground/validator.v10"

	_ "github.com/go-sql-driver/mysql"
)

// Conn コネクション
type Conn struct {
	Db *sql.DB
}

// Data やりとりするデータ
type Data struct {
	Group   string `label:"group"`
	Name    string `label:"name" validate:"required"`
	Message string `label:"message" validate:"required"`
}

// Responce データベースからの戻り
type Responce struct {
	Name    string `label:"name"`
	Message string `label:"message"`
}

var db Conn
var err error

func handler(w http.ResponseWriter, r *http.Request) {
	group := r.FormValue("group")
	// テンプレートをパース
	tpl := template.Must(template.ParseFiles("form.html"))

	resp, err := db.findByGroup(group)
	if err != nil {
		log.Println(err)
	}
	// テンプレートを描画
	tpl.Execute(w, resp)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	type (
		Respose struct {
			Status bool `json:"status"`
		}
	)
	group := r.FormValue("group")
	name := r.FormValue("name")
	message := r.FormValue("message")

	req := Data{
		Group:   group,
		Name:    name,
		Message: message,
	}
	err := validator.New().Struct(&req)
	if err != nil {
		log.Println("error:必須項目がないよ！")
		return
	}
	err = db.insert(req)
	if err != nil {
		log.Println("error:データベースが更新できませんでした")
		return
	}

	resp := Respose{
		Status: true,
	}

	res, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", handler)
	route.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8080", route)
}

// SQL実行
func (db Conn) findByGroup(group string) (responce []Responce, err error) {
	mess := Responce{}

	db, err = db.conn()
	defer db.Db.Close()

	rows, err := db.Db.Query("SELECT `name`, `message` FROM message WHERE `group` = ?", group)
	if err != nil {
		log.Println(err)
	}

	// ポインタに入れる
	for rows.Next() {
		if err = rows.Scan(&mess.Name, &mess.Message); err != nil {
			log.Println(err)
		}
		responce = append(responce, mess)
	}
	return
}

// insert 登録
func (db Conn) insert(req Data) (err error) {
	db, err = db.conn()
	defer db.Db.Close()
	insert, err := db.Db.Prepare("INSERT INTO message(`group`, `name`, `message`, `create_at`) VALUES(?, ?,?,?)")
	if err != nil {
		return err
	}

	_, err = insert.Exec(
		req.Group,
		req.Name,
		req.Message,
		time.Now(),
	)
	return err
}

// conn コネクションプール
func (c Conn) conn() (db Conn, err error) {
	c.Db, err = sql.Open("mysql", "sendaigo:&5Y5nVDs@tcp(35.226.16.11:3306)/handson?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal("db error.")
	}
	db = c
	return
}
