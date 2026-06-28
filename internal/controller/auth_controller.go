package controller

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func registrationHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	defer c.Request.Body.Close()

	if err != nil {
		panic(err)
	}

	userData := strings.Split(string(body), "#")

	//save user
	fmt.Println(userData)

	c.Status(http.StatusOK)
}

func loginHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	defer c.Request.Body.Close()

	if err != nil {
		panic(err)
	}

	credentialString := strings.Split(string(body), "#")

	//check cred

	//return cookie

	c.String(http.StatusOK, credentialString[0])
}
