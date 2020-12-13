package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ShortURL struct
type ShortURL struct {
	Hash        string             `bson:"hash" json:"hash,omitempty"`
	OriginalURL string             `bson:"original_url" json:"original_url"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id,omitempty"`
	ExpiresAt   int                `bson:"expires_at" json:"expires_at,omitempty"`
	CreatedAt   int64              `bson:"created_at" json:"created_at,omitempty"`
}
