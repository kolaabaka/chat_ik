package repository

import (
	"chat_ik/internal/config"
	"chat_ik/internal/repository/entity"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"
)

type UserRepository struct {
	config *config.Config
}

func (r *UserRepository) initDB() (*sql.DB, error) {
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

func NewSqlLiteUserRepository(config config.Config) *UserRepository {
	return &UserRepository{config: &config}
}

func (r *UserRepository) SaveUser(user entity.User) {
	con, err := r.initDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	hash := sha256.Sum256([]byte(user.Hash))

	_, err = con.Exec("INSERT INTO \"users\"(nickname, hash) VALUES(?, ?);", user.Nickname, fmt.Sprintf("%x", hash))
	if err != nil {
		panic(err)
	}

}

func (r *UserRepository) LoginUser(user entity.User) string {
	con, err := r.initDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	hash := sha256.Sum256([]byte(user.Hash))
	row, err := con.Query("SELECT id FROM \"users\" WHERE nickname = ? AND hash = ?;", user.Nickname, fmt.Sprintf("%x", hash))

	if err != nil || !row.Next() {
		return ""
	}

	defer row.Close()

	bufSecret := make([]byte, 32)

	rand.Read(bufSecret)

	//todo: save sessions

	return string(bufSecret)
}
