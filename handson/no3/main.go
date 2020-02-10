package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Conn コネクション
type Conn struct {
	Db *sql.DB
}
type Request struct {
	Group   string
	Name    string
	Message string
}

var db Conn
var err error

func handler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	tpl := template.Must(template.ParseFiles("form.html"))

	m := map[string]string{
		"group":   "1",
		"name":    "名前",
		"message": "メッセージ",
	}
	// テンプレートを描画
	tpl.Execute(w, m)
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

	req := Request{
		Group:   group,
		Name:    name,
		Message: message,
	}

	err = db.insert(req)

	log.Printf("error: %v", err)
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

// insert 登録
func (db Conn) insert(req Request) (err error) {
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
