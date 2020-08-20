package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Conn コネクション
type Conn struct {
	Db *sql.DB
}

// Request グループIDを取得する
type Request struct {
	Group string
}

// Response データベースからの戻り
type Response struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// コネクション
var db Conn

// エラー
var err error

func handler(w http.ResponseWriter, r *http.Request) {
	req := Request{
		Group: "1",
	}

	resp, err := db.findByGroup(req.Group)

	if err != nil {
		log.Println(err)
	}
	// JSONの生成
	res, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", handler)
	http.ListenAndServe(":8080", route)
}

// SQL実行
func (db Conn) findByGroup(group string) (response []Response, err error) {
	mess := Response{}

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
		response = append(response, mess)
	}
	return
}

// conn コネクションプールするレシーバ
func (c Conn) conn() (db Conn, err error) {
	c.Db, err = sql.Open("mysql", "sendaigo:&5Y5nVDs@tcp(35.226.16.11:3306)/handson?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal("db error.")
	}
	db = c
	return
}
