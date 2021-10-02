package main

import (
	"fmt"
	"net/http"
)

func index( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Hello World</h1>")
}
func check_healthy( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<h1>Check Healthy</h1>")
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/check_healthy",check_healthy)
	fmt.Println("Server port 3000 starting ....")
	http.ListenAndServe(":3000",nil)
}