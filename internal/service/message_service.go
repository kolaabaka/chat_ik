package service

import (
	"chat_ik/internal/repository"
	"chat_ik/internal/repository/entity"
)

type MessageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repository repository.MessageRepository) *MessageService {
	return &MessageService{repo: repository}
}

func (s *MessageService) GetAllOppenents(author string) []string {
	return s.repo.GetAllOppenents(author)
}

func (s *MessageService) GetDialog(author string, opponent string) []entity.Message {
	return s.repo.GetDialog(author, opponent)
}

func (s *MessageService) AddMessage(message entity.Message) (bool, error) {
	return s.repo.AddMessage(message)
}
