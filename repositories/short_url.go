package repositories

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/models"
)

// ShortURLRepository abstraction
type ShortURLRepository struct {
	Database *mongo.Database
	Cache    *redis.Client
}

// NewShortURLRepository creates an repository with each database layer
func NewShortURLRepository(database *mongo.Database, cache *redis.Client) *ShortURLRepository {
	return &ShortURLRepository{
		Database: database,
		Cache:    cache,
	}
}

// Create method persists an URL object into the database
func (repository *ShortURLRepository) Create(url *models.ShortURL) (*models.ShortURL, error) {
	logger := log.GetLogger()

	collection := repository.Database.Collection("urls")

	insertResult, err := collection.InsertOne(context.TODO(), url)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to store URL on collection")

		return nil, err
	}

	objectID, _ := insertResult.InsertedID.(primitive.ObjectID)

	objectIDAsString := objectID.String()

	logger.WithFields(logrus.Fields{
		"url_object_id": objectIDAsString,
	}).Info("Successfully stored URL on mongoDB")

	return url, nil
}

// FindURLByFilter method returns the URL from mongoDB collection by some filter
func (repository *ShortURLRepository) FindURLByFilter(filter bson.M) (*models.ShortURL, error) {
	logger := log.GetLogger()

	collection := repository.Database.Collection("urls")

	var shortURL models.ShortURL

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := collection.FindOne(
		ctx,
		filter,
	).Decode(&shortURL)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Debug("URL not found for filter")

		return nil, err
	}

	logger.Info("Successfully found URL")

	return &shortURL, nil
}

// SaveToCache method persists an URL object into the cache layer
func (repository *ShortURLRepository) SaveToCache(url *models.ShortURL) (*models.ShortURL, error) {
	logger := log.GetLogger()

	ctx := context.Background()
	key := url.Hash
	value := url.OriginalURL
	expiration := time.Minute * time.Duration(url.ExpiresAt)

	err := repository.Cache.Set(ctx, key, value, expiration).Err()

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to cache URL on redis")

		return nil, err
	}

	logger.Info("Successfully cached URL on redis")

	return url, nil
}

// FindOnCacheByHash method finds an URL on redis
func (repository *ShortURLRepository) FindOnCacheByHash(hash string) (string, error) {
	logger := log.GetLogger()

	ctx := context.Background()

	originalURL, err := repository.Cache.Get(ctx, hash).Result()

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to find URL on redis")

		return "", err
	}

	logger.Info("Found URL on cache")

	return originalURL, nil
}
