package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func name(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
	fmt.Fprintf(res, req.FormValue("name"))
}

func main() {
	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
}
