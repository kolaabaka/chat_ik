package app

import (
	"chat_ik/internal/config"
	"chat_ik/internal/controller"
	sqlite "chat_ik/internal/repository/sqlite"
	"chat_ik/internal/service"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(done chan os.Signal) {
	//todo: make syncronised initDb method
	cfg := config.MustInit()

	repoMessage := sqlite.NewSqlLiteMessageRepository(cfg)
	repoSession := sqlite.NewSqlLiteSessionRepository(cfg)
	repoUser := sqlite.NewSqlLiteUserRepository(cfg, repoSession)

	servMessage := service.NewMessageService(repoMessage)
	servUser := service.NewUserService(repoUser)

	fmt.Print(servMessage)
	fmt.Print(repoSession)

	userController := controller.NewUserController(servUser)

	r := gin.Default()
	controller.SetupRoutes(r, userController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go server.ListenAndServe()

	context := context.TODO()

	<-done
	fmt.Println("SERVER was interrupted")

	server.Shutdown(context)
}
