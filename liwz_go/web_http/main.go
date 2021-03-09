package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2.解析模板
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err: %v", err)
		return
	}
	// 3.渲染模板
	name := "小王子"
	err = tmpl.Execute(w, name)
	if err != nil {
		fmt.Println("render template failed, err: %v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9010", nil)
	if err != nil {
		fmt.Println("Http server fialed,err: %v", err)
		return
	}
}
