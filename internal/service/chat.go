package service

import (
	"github.com/MicroMolekula/gpt-service/internal/models"
	"github.com/MicroMolekula/gpt-service/internal/repository"
	"time"
)

type ChatService struct {
	messageRepository *repository.MessageRepository
	gptService        *GptService
}

func NewChatService(gptService *GptService, messageRepository *repository.MessageRepository) *ChatService {
	return &ChatService{gptService: gptService, messageRepository: messageRepository}
}

func (s *ChatService) SendMessage(message string, user *models.User) (*models.Message, error) {
	userMessage := &models.Message{
		UserId:  user.ID,
		Date:    time.Now(),
		Type:    false,
		Context: message,
	}
	gptResp, err := s.gptService.QueryLite("Ты фитнес-тренер, способный помочь пользователю в вопросах по фитнесу. Общайся кратко и на простом языке. Отвечай только по теме.", message)
	if err != nil {
		return nil, err
	}
	chatMessage := &models.Message{
		UserId:  user.ID,
		Date:    time.Now(),
		Type:    true,
		Context: gptResp.Alternatives[0].GptMessage.Text,
	}
	if err = s.messageRepository.SaveThoMessage(userMessage, chatMessage); err != nil {
		return nil, err
	}
	return chatMessage, nil
}

func (s *ChatService) GetAllMessages(user *models.User) ([]*models.Message, error) {
	messages, err := s.messageRepository.FindByUserId(user.ID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
