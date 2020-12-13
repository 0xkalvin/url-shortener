package schemas

// ShortURLPostSchema struct
type ShortURLPostSchema struct {
	OriginalURL string `json:"original_url" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	ExpiresAt   int    `json:"expires_at,omitempty"`
}
