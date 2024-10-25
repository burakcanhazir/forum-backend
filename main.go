package main

import (
	"burakforum/controllers"
	"burakforum/database"
	"burakforum/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()

	// CORS ayarlarını yapın
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	protected := r.PathPrefix("/api/v1").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// middleware gerektirmeyen endpointler USER REGİSTER VE LOGİN
	r.HandleFunc("/api/v1/register", controllers.Register).Methods("POST") // YENİ KULLANICI OLUŞTUR   --YAPILDI
	r.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")       // GİRİŞ YAP                --YAPILDI
	r.HandleFunc("/api/v1/logout", controllers.Logout).Methods("POST")     // çıkış yapmak             --YAPILDI

	// USERS İŞLEMLERİ
	protected.HandleFunc("/getusers", controllers.GetUsers).Methods("GET")          // KULLANICILARI GÖRÜNTÜLE
	protected.HandleFunc("/getusers/{id}", controllers.GetUsersID).Methods("GET")   // X KULLANICIYI GÖRÜNTÜLE
	protected.HandleFunc("/deleteusers", controllers.DeleteUsers).Methods("DELETE") // KULLANICI SİLME

	// middleware gerektirmeyen endpointler POST
	r.HandleFunc("/api/v1/homepage", controllers.GetPosts).Methods("GET") // tüm postları görüntüle         --YAPILDI

	// POST İŞLEMLERİ
	protected.HandleFunc("/createpost", controllers.CreatePost).Methods("POST")             // YENİ POST OLUŞTURMA & CATEGORİ İÇİNDE             --YAPILDI
	r.HandleFunc("/api/v1/getpost/{id}", controllers.GetPostID).Methods("GET")              // belirli gönderiyi görüntüler                      --YAPILDI
	protected.HandleFunc("/users/getpost/{id}", controllers.GetUsersPostsID).Methods("GET") // belirli kullanıcının tüm gönderilerini görüntüle   --YAPILDI
	protected.HandleFunc("/deletepost/{id}", controllers.DeletePost).Methods("DELETE")      // POST SİLME

	// LİKE İŞLEMLERİ
	protected.HandleFunc("/{post_id}/like", controllers.LikePost).Methods("POST") // POSTA LİKE ATMA                                        --YAPILDI
	protected.HandleFunc("/mylikes", controllers.UsersLikesPost).Methods("GET")   // like attıklarını görüntüleme                            --YAPILDI

	// commit işlemleri
	protected.HandleFunc("/createcommit/{id}", controllers.CreateCommit).Methods("POST")
	protected.HandleFunc("/deletecommit/{postID}/{commitID}", controllers.DeleteCommit).Methods("DELETE")

	// Kategori
	r.HandleFunc("/api/v1/getcategoriespost/{id}", controllers.GetCategoriesPost).Methods("GET") // --YAPILDI

	// CORS middleware'ini uygulayın
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
