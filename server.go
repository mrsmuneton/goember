package main

import(
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func KittensHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"kittens": [{"id": 1, "name": "Wally", "picture": "http://placekitten.com/200/200"},{"id": 2, "name": "Ray", "picture": "http://placekitten.com/300/300"}]}`))
}

func main(){
	log.Println("Start server")

	log.Println("setting mux r router...")
	r := mux.NewRouter()
	r.HandleFunc("/api/kittens", KittensHandler).Methods("GET")
	http.Handle("/api/", r)

	log.Println("setting fileserver for .public/index.html at /")
	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Println("listening to port 8080")
	http.ListenAndServe(":8080",nil)
}
