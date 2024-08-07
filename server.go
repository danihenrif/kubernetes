package main

import "net/http"
func main(){
	http.HandleFunc("/",Hello)
	http.ListenAndServe(":8070", nil)	
}

func Hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello Full Cycle </h1>"))
} 