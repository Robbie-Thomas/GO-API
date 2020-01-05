package main

import(
	"github.com/gorilla/mux"
	"new/http"
	"encoding/json"
)

type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`

}

var posts []POST

func main() {
 router := mux.NewRouter()

 router.HandleFunc("/posts", getPosts).Methods("GET")
 router.HandleFunc("/posts", createPost).Methods("POST")
 router.HandleFunc("/posts/{id}", getPost).Methods("GET")
 router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
 router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
 http.ListenAndServer(":8000", router)
}