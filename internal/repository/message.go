package repository

import (
	"github.com/MicroMolekula/gpt-service/internal/models"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (m *MessageRepository) Save(message *models.Message) error {
	return m.db.Create(message).Error
}

func (m *MessageRepository) SaveThoMessage(userMessage *models.Message, chatMessage *models.Message) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(userMessage).Error; err != nil {
			return err
		}
		if err := tx.Create(chatMessage).Error; err != nil {
			return err
		}
		return nil
	})
}

func (m *MessageRepository) FindByUserId(userId uint) ([]*models.Message, error) {
	var message []*models.Message
	if err := m.db.Where("user_id = ?", userId).Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}
