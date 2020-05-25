package main

import (
	"courseSelection/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.PingPongHandler)
	http.HandleFunc("/course-selection", handler.Handler)
	http.HandleFunc("/course-selections", handler.AllHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
