package main

import (
	"chat_ik/internal/config"
	sqlite "chat_ik/internal/repository/sqlite"
	"chat_ik/internal/service"
	"fmt"
)

func main() {
	cfg := config.MustInit()
	repo := sqlite.NewSqlLiteRepository(cfg)
	serv := service.NewService(repo)

	fmt.Println(serv)
}
