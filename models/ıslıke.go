package models

type LikeDislike struct {
	ID     string `json:"id"`      // Like/Dislike ID'si
	PostID string `json:"post_id"` // Beğeni/beğenmeme yapılan postun ID'si
	UserID string `json:"user_id"` // Beğeni/beğenmeme yapan kullanıcının ID'si
	IsLike bool   `json:"is_like"` // Beğeni mi beğenmeme mi olduğunu belirten bayrak
}
