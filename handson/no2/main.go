package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Conn コネクション
type Conn struct {
	Db *sql.DB
}
type Request struct {
	Group string
}

type Responce struct {
	Message string
}

var db Conn
var err error

func handler(w http.ResponseWriter, r *http.Request) {
	type (
		Respose struct {
			Status bool `json:"status"`
		}
	)

	req := Request{
		Group: "1",
	}

	re, err := db.findByGroup(req.Group)

	fmt.Printf("log: %v", re)
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
	http.ListenAndServe(":8080", route)
}

// SELECT
func (db Conn) findByGroup(group string) (responce []Responce, err error) {
	mess := Responce{}

	db, err = db.conn()
	defer db.Db.Close()

	rows, err := db.Db.Query("SELECT `message` FROM message WHERE `group` = ?", group)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err = rows.Scan(&mess.Message); err != nil {
			log.Println(err)
		}
		responce = append(responce, mess)
	}

	return
}

// conn コネクションプールする、レシーバ
func (c Conn) conn() (db Conn, err error) {
	c.Db, err = sql.Open("mysql", "sendaigo:&5Y5nVDs@tcp(35.226.16.11:3306)/handson?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal("db error.")
	}
	db = c
	return
}
