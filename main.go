package main

import(
	"github.com/gorilla/mux"
	"new/http"
	"encoding/json"
	"math/rand"
	"strconv"
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

func getPosts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Post{})
}


func createPost(w http.ResponseWriter, r *http.Request){
	
}