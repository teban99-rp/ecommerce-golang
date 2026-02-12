package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/models"
	"github.com/teban99-rp/ecommerce-golang/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, _ := c.service.GetUsers()
	ctx.JSON(http.StatusOK, users)
}
