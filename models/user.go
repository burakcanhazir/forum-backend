package models

// User struct defines the structure for user data
type User struct {
	ID    string `json:"id"`    // Kullanıcı ID'si
	Name  string `json:"name"`  // Kullanıcı adı
	Email string `json:"email"` // Kullanıcı email'i
}
