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
	cfg := config.MustInit()
	repo := sqlite.NewSqlLiteRepository(cfg)
	serv := service.NewService(repo)

	fmt.Println(serv)

	r := gin.Default()
	controller.SetupRoutes(r)

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
