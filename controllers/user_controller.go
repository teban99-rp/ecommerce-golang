package controllers

import (
	"net/http"
	"strconv"

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

// Services
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

// Views
func (c *UserController) CreateUserView(ctx *gin.Context) {
	var user models.User

	name := ctx.PostForm("name")
	lastName := ctx.PostForm("lastname")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	user.Name = name
	user.LastName = lastName
	user.Email = email
	user.Password = password

	if err := c.service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Redirect(http.StatusSeeOther, "/view/register")
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/view/products")
}

func (c *UserController) GetUsersView(ctx *gin.Context) {
	users, _ := c.service.GetUsers()

	ctx.HTML(http.StatusOK, "layout", gin.H{
		"title": "Inicio",
		"view":  "home",
		"users": users,
	})
}

func (c *UserController) LoginView(ctx *gin.Context) {

	var input dto.LoginDTO

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	input.Email = email
	input.Password = password

	token, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.HTML(http.StatusUnauthorized, "layout", gin.H{
			"title":     "Iniciar Sesión",
			"view":      "login",
			"error":     "Correo o contraseña incorrectos. Inténtalo de nuevo.", // El mensaje
			"logged_in": false,
		})
		return
	}

	ctx.SetCookie(
		"token",
		token,
		86400, // 1 día
		"/view",
		"localhost",
		false,
		true,
	)

	user, err := c.service.FindByEmail(input.Email)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/view/products")
		return
	}

	ctx.SetCookie(
		"user_id",
		strconv.Itoa(int(user.ID)),
		86400, // 1 día
		"/view",
		"localhost",
		false,
		true,
	)

	ctx.SetCookie(
		"role",
		user.Role,
		86400, // 1 día
		"/view",
		"localhost",
		false,
		true,
	)

	if user.Role == "admin" {
		ctx.Redirect(http.StatusSeeOther, "/view/admin/dashboard")
	} else {
		ctx.Redirect(http.StatusSeeOther, "/view/products")
	}
}

func (c *UserController) LogoutView(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("token", "", -1, "/view", "localhost", false, true)
	ctx.SetCookie("user_id", "", -1, "/view", "localhost", false, true)
	ctx.SetCookie("role", "", -1, "/view", "localhost", false, true)
	ctx.Redirect(http.StatusSeeOther, "/view/")
}
