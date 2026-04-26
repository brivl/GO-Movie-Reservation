package handlers

import (
	"movies-app/dtos"
	"movies-app/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepo *repositories.UserRepository
}

func (handler *UserHandler) Register(c *gin.Context) {
	var req dtos.UserCreateDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if existingUser, _ := handler.UserRepo.GetByEmail(req.Email); existingUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
	}

	if err := handler.UserRepo.Create(req.Email, req.Password, req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (handler *UserHandler) Login(c *gin.Context) {
}

func (handler *UserHandler) SetAdmin(c *gin.Context) {
	// check if request has an admin rights to set the user to admin
	// check if he's authorized
}
