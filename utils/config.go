package utils

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	// Genel ayarlar
	AppEnv  string
	AppPort string
	AppHost string

	// Veritabanı ayarları
	DbDriver string
	DbPath   string

	// JWT ayarları
	JwtSecret     string
	JwtExpiration string

	// CORS ayarları
	AllowedOrigins string
	AllowedMethods []string
	AllowedHeaders []string

	// Rate limiting
	RateLimit  string
	RateWindow string
)

func Init() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// .env dosyasındaki değişkenleri oku
	AppEnv = os.Getenv("APP_ENV")
	AppPort = os.Getenv("APP_PORT")
	AppHost = os.Getenv("APP_HOST")

	DbDriver = os.Getenv("DB_DRIVER")
	DbPath = os.Getenv("DB_PATH")

	JwtSecret = os.Getenv("JWT_SECRET")
	JwtExpiration = os.Getenv("JWT_EXPIRATION")

	AllowedOrigins = os.Getenv("ALLOWED_ORIGINS")

	AllowedMethods = strings.Split(os.Getenv("ALLOWED_METHODS"), ",")

	AllowedHeaders = strings.Split(os.Getenv("ALLOWED_HEADERS"), ",")

	RateLimit = os.Getenv("RATE_LIMIT")
	RateWindow = os.Getenv("RATE_LIMIT_WINDOW")
}
