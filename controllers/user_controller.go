package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teban99-rp/ecommerce-golang/dto"
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

	name := ctx.PostForm("name")
	lastName := ctx.PostForm("lastname")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	user.Name = name
	user.LastName = lastName
	user.Email = email
	user.Password = password

	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := c.service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/products")
	//ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, _ := c.service.GetUsers()
	ctx.JSON(http.StatusOK, users)
}

// Login maneja la autenticación de usuarios y genera un token JWT
func (c *UserController) Login(ctx *gin.Context) {

	var input dto.LoginDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
