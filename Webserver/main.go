package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method !="GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!= nil {
		fmt.Fprintf(w, "ParseForm() err: %v",err)
		return 
	}
	fmt.Fprint(w, "Post Request Successful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =  %v",name)
	fmt.Fprintf(w,"Address = %v",address)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}

}