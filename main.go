package main

import (
	"log"
	"net/http"

	"burakforum/controllers"
	"burakforum/database"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()

	// USERS İŞLEMLERİ
	r.HandleFunc("/api/v1/getusers", controllers.GetUsers).Methods("GET")          // KULLANICILARI GÖRÜNTÜLE
	r.HandleFunc("/api/v1/getusers/{id}", controllers.GetUsersID).Methods("GET")   // X KULLANICIYI GÖRÜNTÜLE
	r.HandleFunc("/api/v1/createusers", controllers.CreateUsers).Methods("POST")   // YENİ KULLANICI OLUŞTUR
	r.HandleFunc("/api/v1/deleteusers", controllers.DeleteUsers).Methods("DELETE") // KULLANICI SİLME

	// POST İŞLEMLERİ
	r.HandleFunc("/api/v1/homepage", controllers.GetPosts).Methods("GET") // tüm postları görüntüle
	// r.HandleFunc("/api/v1/getpostcategories/{id}", controllers.GetPostCategories).Methods("GET") // belirli gönderiyi görüntüler ... // bu kısım sorunlu ; bir post birden fazla kategori barındıracak.
	r.HandleFunc("/api/v1/createpost", controllers.CreatePost).Methods("POST")             // YENİ POST OLUŞTURMA
	r.HandleFunc("/api/v1/getpost/{id}", controllers.GetPostID).Methods("GET")             // belirli gönderiyi görüntüler
	r.HandleFunc("/api/v1/users/{id}/getpost", controllers.GetUsersPostsID).Methods("GET") // belirli kullanıcının tüm gönderilerini görüntüle
	r.HandleFunc("/api/v1/deletepost/{id}", controllers.DeletePost).Methods("DELETE")      // POST SİLME

	log.Fatal(http.ListenAndServe(":8000", r))
}
