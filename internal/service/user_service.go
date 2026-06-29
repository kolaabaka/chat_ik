package service

import (
	"chat_ik/internal/repository"
	"chat_ik/internal/repository/entity"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repo: repository}
}

func (u *UserService) SaveUser(user entity.User) {
	u.repo.SaveUser(user)
}

func (u *UserService) LoginUser(user entity.User) string {
	return u.repo.LoginUser(user)
}
