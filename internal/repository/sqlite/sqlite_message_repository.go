package repository

import (
	"chat_ik/internal/config"
	"chat_ik/internal/repository/entity"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type MessageRepository struct {
	config *config.Config
}

func NewSqlLiteMessageRepository(config config.Config) *MessageRepository {
	return &MessageRepository{config: &config}
}

func (r *MessageRepository) initDB() (*sql.DB, error) {
	db, err := sql.Open(r.config.DataSource.Sql.Type, r.config.DataSource.Sql.URL)
	if err != nil {
		return nil, fmt.Errorf("failed open db: %w", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Second * 10)

	// check connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed ping db: %w", err)
	}

	return db, nil
}

func (r *MessageRepository) GetAllOppenents(author string) []string {
	con, err := r.initDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	return []string{"mock"}
}

func (r *MessageRepository) GetDialog(author string, opponent string) []entity.Message {
	con, err := r.initDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()
	return []entity.Message{}
}

func (r *MessageRepository) AddMessage(entity.Message) (bool, error) {
	con, err := r.initDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()
	return true, nil
}
