package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	tpl := template.Must(template.ParseFiles("index.html"))

	m := map[string]string{
		"name": "Gopaher",
	}
	// テンプレートを描画
	tpl.Execute(w, m)
}

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", handler)
	http.ListenAndServe(":8080", route)
}
