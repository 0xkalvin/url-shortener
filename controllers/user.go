package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/0xkalvin/url-shortener/logger"
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
			gin.H{"error_type": "Invalid request body"},
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
