package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("method:", method) //获取请求的方法
	if method == "GET" {
		t, _ := template.ParseFiles("pages/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Print("key:", k)
		fmt.Println(" val:", strings.Join(v, ""))
	}

	fmt.Fprint(w, "Hello Andaicheng")
}
