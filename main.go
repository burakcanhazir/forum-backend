package main

import (
	"log"
	"net/http"

	"burakforum/controllers"
	"burakforum/database"
	"burakforum/middleware"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()

	protected := r.PathPrefix("/api/v1").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// middleware gerektirmeyen endpointler USER
	r.HandleFunc("/api/v1/register", controllers.Register).Methods("POST") // YENİ KULLANICI OLUŞTUR
	r.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")       // GİRİŞ YAP

	// USERS İŞLEMLERİ
	protected.HandleFunc("/getusers", controllers.GetUsers).Methods("GET")          // KULLANICILARI GÖRÜNTÜLE
	protected.HandleFunc("/getusers/{id}", controllers.GetUsersID).Methods("GET")   // X KULLANICIYI GÖRÜNTÜLE
	protected.HandleFunc("/deleteusers", controllers.DeleteUsers).Methods("DELETE") // KULLANICI SİLME

	// middleware gerektirmeyen endpointler POST
	r.HandleFunc("/api/v1/homepage", controllers.GetPosts).Methods("GET") // tüm postları görüntüle

	// POST İŞLEMLERİ

	// r.HandleFunc("/api/v1/getpostcategories/{id}", controllers.GetPostCategories).Methods("GET") // belirli gönderiyi görüntüler ... // bu kısım sorunlu ; bir post birden fazla kategori barındıracak.
	protected.HandleFunc("/createpost", controllers.CreatePost).Methods("POST")             // YENİ POST OLUŞTURMA
	protected.HandleFunc("/getpost/{id}", controllers.GetPostID).Methods("GET")             // belirli gönderiyi görüntüler
	protected.HandleFunc("/users/{id}/getpost", controllers.GetUsersPostsID).Methods("GET") // belirli kullanıcının tüm gönderilerini görüntüle
	protected.HandleFunc("/deletepost/{id}", controllers.DeletePost).Methods("DELETE")      // POST SİLME

	// LİKE İŞLEMLERİ
	protected.HandleFunc("/{post_id}/like", controllers.LikePost).Methods("POST") // POSTA LİKE ATMA
	protected.HandleFunc("/mylikes", controllers.UsersLikesPost).Methods("GET")   // like attıklarını görüntüleme

	// commit işlemleri
	protected.HandleFunc("/createcommit/{id}", controllers.CreateCommit).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
