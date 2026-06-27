package repository

import entity "chat_ik/internal/repository/entity"

type MessageRepository interface {
	GetAllOppenents(author string) []string
	GetDialog(author string, opponent string) []entity.Message
	AddMessage(entity.Message) (bool, error)
}
