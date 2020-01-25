package models



type URL struct {
	Long string `bson:"long_url" json:"long_url" binding:"required"`
	Short string `bson:"short_url" json:"short_url" binding:"required"`
	CreatedAt int64 `bson:"created_at" json:"created_at"`
}


