.PHONY: mig-up run
mig-up:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir migration sqlite3 ./resources/db/messager.db up

run:
	go run cmd/main.go