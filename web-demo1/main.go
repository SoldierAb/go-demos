package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type  MyMux struct {
}



func (m *MyMux) ServerHttp(w http.ResponseWriter,r *http.Request){
	if r.URL.Path == "/"{
		sayHello(w,r)
		return
	}
	http.NotFound(w,r)
	return
}


func sayHello(w http.ResponseWriter,r *http.Request){
	r.ParseForm();     //参数解析
	scheme := r.URL.Scheme
	path := r.URL.Path

	fmt.Println("scheme; path ",scheme,path)

	for key,val := range r.Form{
		fmt.Println("key - val ",key,strings.Join(val,""))
	}

	fmt.Fprint(w,"hello chengj")
}


func login(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Method)
	if r.Method == "GET"{
		t,_ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	}else{
		fmt.Println("userName",r.Form["username"])
		fmt.Println("password",r.Form["password"])
	}
}


func main(){
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/login", login)
	//mux := &MyMux{}
	//
	err := http.ListenAndServe(":8080",nil)

	if err !=nil{
		log.Fatal("ListenAndServe:   ",err)
	}
}




