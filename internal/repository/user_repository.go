package repository

import "chat_ik/internal/repository/entity"

type UserRepository interface {
	SaveUser(user entity.User)
	LoginUser(user entity.User) string
}
