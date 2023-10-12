package controller

import (
	"net/http"

	"bank-ina-assessment/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginUserController(c *gin.Context) error
}

type authController struct {
	userUsecase usecase.UserUsecase
}

func NewAuthController(userUsecase usecase.UserUsecase) *authController {
	return &authController{
		userUsecase,
	}
}

func (a *authController) LoginUserController(c *gin.Context) {
	code := c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	userData, err := a.userUsecase.LoginUser(c, code, pathUrl)
	if err.NotNil() {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   userData,
	})

}
