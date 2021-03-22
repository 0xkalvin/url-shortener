package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/0xkalvin/url-shortener/logger"
	"github.com/0xkalvin/url-shortener/repositories"
	"github.com/0xkalvin/url-shortener/schemas"
	"github.com/0xkalvin/url-shortener/services"
)

// UserController struct
type UserController struct {
	Service services.UserService
}

// NewUserController creates an controller layer
func NewUserController(service services.UserService) *UserController {
	return &UserController{
		Service: service,
	}
}

// Create an user
func (controller UserController) Create(context *gin.Context) {
	var userPayload schemas.UserPostSchema

	err := context.Bind(&userPayload)

	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"error_type":    "Invalid request body",
				"error_message": err.Error(),
			},
		)

		return
	}

	user, err := controller.Service.CreateUser(userPayload.Name, userPayload.Email)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error_type": "Internal server error"},
		)

		return
	}
	logger := log.GetLogger()

	logger.Info("Finished to create user")

	context.JSON(http.StatusOK, user)
}

// Show returns an user
func (controller UserController) Show(context *gin.Context) {
	logger := log.GetLogger()

	userIDAsString := context.Param("id")

	userObjectID, err := primitive.ObjectIDFromHex(userIDAsString)

	if err != nil {
		logger.Error("Failed to build object id for user")

		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error_type":    "Bad request",
				"error_message": "Invalid user ID",
			},
		)

		return
	}

	user, err := controller.Service.FindOneUser(userObjectID)

	if err != nil {
		switch err {
		case repositories.ErrUserNotFound:
			context.JSON(
				http.StatusNotFound,
				gin.H{
					"error_message": fmt.Sprintf("User %s not found", userIDAsString),
					"error_type":    "Not found",
				},
			)

			return
		default:
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"error_type": "Internal server error"},
			)

			return
		}
	}

	context.JSON(http.StatusOK, user)
}
