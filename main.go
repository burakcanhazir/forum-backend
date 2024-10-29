package main

import (
	"log"
	"net/http"
	"time"

	"burakforum/controllers"
	"burakforum/database"
	"burakforum/middleware"

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
	r.Use(middleware.RateLimitMiddleware)
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
	protected.HandleFunc("/deletepost/{id}", controllers.DeletePost).Methods("DELETE")      // POST SİLME            								---YAPILDI

	// LİKE İŞLEMLERİ
	protected.HandleFunc("/{post_id}/like", controllers.LikePost).Methods("POST") // POSTA LİKE ATMA                                        --YAPILDI
	protected.HandleFunc("/mylikes", controllers.UsersLikesPost).Methods("GET")   // like attıklarını görüntüleme                            --YAPILDI

	// commit işlemleri
	protected.HandleFunc("/createcommit/{id}", controllers.CreateCommit).Methods("POST") // yapıldı
	protected.HandleFunc("/deletecommit/{postID}/{commitID}", controllers.DeleteCommit).Methods("DELETE")
	protected.HandleFunc("/getcommits/{id}", controllers.GetCommits).Methods("GET")

	// Kategori
	r.HandleFunc("/api/v1/getcategoriespost/{id}", controllers.GetCategoriesPost).Methods("GET") // --YAPILDI

	// CORS middleware'ini uygulayın
	handler := c.Handler(r)
	// Sunucu ayarları ile birlikte başlat
	server := &http.Server{
		Addr:              ":8000",
		Handler:           handler,
		ReadTimeout:       2 * time.Second,   // İsteği okumak için maksimum süre
		WriteTimeout:      10 * time.Second,  // Yanıtı istemciye göndermek için maksimum süre
		IdleTimeout:       120 * time.Second, // Boşta bekleme süresi
		ReadHeaderTimeout: 5 * time.Second,   // Başlıkları okumak için maksimum süre
	}

	log.Println("Sunucu başlatılıyor...")
	log.Fatal(server.ListenAndServe())
}
