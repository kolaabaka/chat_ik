package service

import (
	"chat_ik/internal/repository"
	"chat_ik/internal/repository/entity"
)

type Service struct {
	repo repository.MessageRepository
}

func NewService(repository repository.MessageRepository) *Service {
	return &Service{repo: repository}
}

func (s *Service) GetAllOppenents(author string) []string {
	return s.repo.GetAllOppenents(author)
}

func (s *Service) GetDialog(author string, opponent string) []entity.Message {
	return s.repo.GetDialog(author, opponent)
}

func (s *Service) AddMessage(message entity.Message) (bool, error) {
	return s.repo.AddMessage(message)
}
