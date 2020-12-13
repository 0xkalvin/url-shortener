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

// FindOriginalURLByHash method returns the original URL from mongoDB collection by hash
func (repository *ShortURLRepository) FindOriginalURLByHash(hash string) (string, error) {
	logger := log.GetLogger()

	collection := repository.Database.Collection("urls")

	var shortURL models.ShortURL

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := collection.FindOne(
		ctx,
		bson.M{"hash": hash},
	).Decode(&shortURL)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to find URL on MongoDB collection")

		return "", err
	}

	logger.Info("Successfully found URL")

	return shortURL.OriginalURL, nil
}

// SaveToCache method persists an URL object into the cache layer
func (repository *ShortURLRepository) SaveToCache(url *models.ShortURL) (*models.ShortURL, error) {
	logger := log.GetLogger()

	ctx := context.Background()
	key := url.Hash
	value := url.OriginalURL
	expiration := time.Second * time.Duration(url.ExpiresAt)

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
