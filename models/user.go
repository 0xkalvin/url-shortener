package models

// User struct
type User struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}
