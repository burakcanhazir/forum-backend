package main

import (
	"log"
	"net/http"
	"time"

	"forumbackend/controllers"
	"forumbackend/database"
	"forumbackend/middleware"
	"forumbackend/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()
	utils.Init()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{utils.AllowedOrigins}, // .env dosyasındaki izinli kökeni kullanıyoruz
		AllowCredentials: true,                           // Credentials (cookie, auth) gönderilecekse true olmalı
		AllowedMethods:   utils.AllowedMethods,           // methods artık dilim olarak alınıyor
		AllowedHeaders:   utils.AllowedHeaders,           // headers artık dilim olarak alınıyor
		ExposedHeaders:   []string{"Authorization"},      // Yanıt başlıklarına ekleyeceğiniz başlıklar
	})
	r.Use(middleware.RateLimitMiddleware)
	protected := r.PathPrefix("/api/v1").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// middleware gerektirmeyen endpointler USER REGİSTER VE LOGİN
	r.HandleFunc("/api/v1/register", controllers.Register).Methods("POST") // YENİ KULLANICI OLUŞTUR
	r.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")       // GİRİŞ YAP
	r.HandleFunc("/api/v1/logout", controllers.Logout).Methods("POST")     // çıkış yapmak

	// USERS İŞLEMLERİ
	protected.HandleFunc("/getusers", controllers.GetUsers).Methods("GET")          // KULLANICILARI GÖRÜNTÜLE
	protected.HandleFunc("/getusers/{id}", controllers.GetUsersID).Methods("GET")   // X KULLANICIYI GÖRÜNTÜLE
	protected.HandleFunc("/deleteusers", controllers.DeleteUsers).Methods("DELETE") // KULLANICI SİLME

	// middleware gerektirmeyen endpointler POST
	r.HandleFunc("/api/v1/homepage", controllers.GetPosts).Methods("GET") // tüm postları görüntüle

	// POST İŞLEMLERİ
	protected.HandleFunc("/createpost", controllers.CreatePost).Methods("POST")             // YENİ POST OLUŞTURMA & CATEGORİ İÇİNDE
	r.HandleFunc("/api/v1/getpost/{id}", controllers.GetPostID).Methods("GET")              // belirli gönderiyi görüntüler
	protected.HandleFunc("/users/getpost/{id}", controllers.GetUsersPostsID).Methods("GET") // belirli kullanıcının tüm gönderilerini görüntüle
	protected.HandleFunc("/deletepost/{id}", controllers.DeletePost).Methods("DELETE")      // POST SİLME

	// LİKE İŞLEMLERİ
	protected.HandleFunc("/{post_id}/like", controllers.LikePost).Methods("POST") // POSTA LİKE ATMA
	protected.HandleFunc("/mylikes", controllers.UsersLikesPost).Methods("GET")   // like attıklarını görüntüleme

	// commit işlemleri
	protected.HandleFunc("/createcommit/{id}", controllers.CreateCommit).Methods("POST")
	protected.HandleFunc("/deletecommit/{postID}/{commitID}", controllers.DeleteCommit).Methods("DELETE")
	protected.HandleFunc("/getcommits/{id}", controllers.GetCommits).Methods("GET")

	// Kategori
	r.HandleFunc("/api/v1/getcategoriespost/{id}", controllers.GetCategoriesPost).Methods("GET")

	// CORS middleware'i entegre ediyorum
	handler := c.Handler(r)
	// sunucuyu başlat
	server := &http.Server{
		Addr:              ":" + utils.AppPort,
		Handler:           handler,
		ReadTimeout:       2 * time.Second,   // İsteği okumak için maksimum süre
		WriteTimeout:      10 * time.Second,  // Yanıtı istemciye göndermek için maksimum süre
		IdleTimeout:       120 * time.Second, // Boşta bekleme süresi
		ReadHeaderTimeout: 5 * time.Second,   // Başlıkları okumak için maksimum süre
	}

	log.Println("Sunucu başlatılıyor...")
	log.Fatal(server.ListenAndServe())
}
