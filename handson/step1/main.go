package main

import (
	"html/template"
	"net/http"
)

var name string

func handler(w http.ResponseWriter, r *http.Request) {
	// 変数
	name := "Gopaher"
	// テンプレートをパース
	tpl := template.Must(template.ParseFiles("index.html"))

	// 配列
	m := map[string]string{
		"name": name,
	}
	// テンプレートを描画
	tpl.Execute(w, m)
}

func main() {
	// ルーティングとハンドラとコンテキスト
	route := http.NewServeMux()
	route.HandleFunc("/", handler)
	http.ListenAndServe(":8080", route)
}
