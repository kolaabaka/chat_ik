package repository

import entity "chat_ik/internal/repository/entity"

type SessionRepository interface {
	AddSession(session entity.Session)
	CheckSession(hash string) entity.Session
}
