package controller

import (
	"chat_ik/internal/repository/entity"
	"chat_ik/internal/service"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: *service}
}

func (uc *UserController) registrationHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	defer c.Request.Body.Close()

	if err != nil {
		panic(err)
	}

	userData := strings.Split(string(body), "#")

	userToSave := entity.User{
		Nickname: userData[0],
		Hash:     userData[1],
	}

	uc.service.SaveUser(userToSave)

	c.Status(http.StatusOK)
}

func (uc *UserController) loginHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	defer c.Request.Body.Close()

	if err != nil {
		panic(err)
	}

	credentialString := strings.Split(string(body), "#")

	loginUser := entity.User{
		Nickname: credentialString[0],
		Hash:     credentialString[1],
	}

	cookie := uc.service.LoginUser(loginUser)

	if cookie == "" {
		c.String(http.StatusForbidden, cookie)

	} else {
		c.String(http.StatusOK, cookie)
	}

}
