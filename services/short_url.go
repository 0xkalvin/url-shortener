package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/models"
	"github.com/0xkalvin/url-shortener/repositories"
	"github.com/0xkalvin/url-shortener/schemas"

	"gopkg.in/lucsky/cuid.v1"
)

// ShortURLService interface
type ShortURLService struct {
	ShortURLRepository repositories.ShortURLRepository
	UserRepository     repositories.UserRepository
}

// NewShortURLService creates an service layer
func NewShortURLService(
	shortURLRepository repositories.ShortURLRepository,
	userRepository repositories.UserRepository) *ShortURLService {

	return &ShortURLService{
		ShortURLRepository: shortURLRepository,
		UserRepository:     userRepository,
	}
}

// CreateURL creates an user
func (service *ShortURLService) CreateURL(payload schemas.ShortURLPostSchema) (*models.ShortURL, error) {
	logger := log.GetLogger()

	userObjectID, err := primitive.ObjectIDFromHex(payload.UserID)

	if err != nil {
		logger.Error("Failed to build user ID for URL")

		return nil, err
	}

	user, err := service.UserRepository.FindByID(userObjectID)

	if err != nil {
		logger.Error("Failed to fetch user entity")

		return nil, err
	}

	filter := bson.M{
		"original_url": payload.OriginalURL,
		"user_id":      userObjectID,
	}

	alreadyExists, err := service.ShortURLRepository.FindURLByFilter(filter)

	if alreadyExists != nil {
		logger.Error("Original URL already exists, returning its object")

		return alreadyExists, nil
	}

	shortURL := &models.ShortURL{
		Hash:        cuid.New(),
		OriginalURL: payload.OriginalURL,
		UserID:      user.ID,
		ExpiresAt:   payload.ExpiresAt,
		CreatedAt:   time.Now().Unix(),
	}

	url, err := service.ShortURLRepository.Create(shortURL)

	if err != nil {
		logger.Error("Failed to create URL entity")

		return nil, err
	}

	_, err = service.ShortURLRepository.SaveToCache(shortURL)

	if err != nil {
		logger.Error("Failed to cache URL entity")

		return nil, err
	}

	return url, nil
}

// FindOneURLByHash retuns an URL if it exists
func (service *ShortURLService) FindOneURLByHash(hash string) (string, error) {
	logger := log.GetLogger()

	originalURL, err := service.ShortURLRepository.FindOnCacheByHash(hash)

	if err == nil {
		return originalURL, nil
	}

	logger.Info("Looking for URL on MongoDB collection")

	filter := bson.M{
		"hash": hash,
	}

	url, err := service.ShortURLRepository.FindURLByFilter(filter)

	if err != nil {
		logger.Error("Failed to find URL entity")

		return "", err
	}

	return url.OriginalURL, nil
}
