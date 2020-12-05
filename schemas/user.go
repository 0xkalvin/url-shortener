package schemas

// UserPostSchema schema for post method payload
type UserPostSchema struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
