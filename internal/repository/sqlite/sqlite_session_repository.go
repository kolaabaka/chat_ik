package repository

import (
	"chat_ik/internal/config"
	"chat_ik/internal/repository/entity"
	"database/sql"
	"fmt"
	"time"
)

type SessionRepository struct {
	config *config.Config
}

func NewSqlLiteSessionRepository(config config.Config) *SessionRepository {
	return &SessionRepository{config: &config}
}

func (rs *SessionRepository) initDB() (*sql.DB, error) {
	db, err := sql.Open(rs.config.DataSource.Sql.Type, rs.config.DataSource.Sql.URL)
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

func (rs *SessionRepository) AddSession(session entity.Session) {
	con, err := rs.initDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()
	_, err = con.Exec("INSERT INTO \"sessions\"(id, hash) VALUES(?, ?);", session.Id, fmt.Sprintf("%x", session.Hash))
	if err != nil {
		panic(err)
	}
}

func (rs *SessionRepository) CheckSession(hash string) entity.Session {
	con, err := rs.initDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	row, err := con.Query("SELECT id FROM \"sessions\" WHERE hash = ?;", fmt.Sprintf("%x", hash))

	session := entity.Session{}

	if err != nil || !row.Next() {
		return session
	}

	defer row.Close()

	row.Scan(&session.Id)

	return session
}
