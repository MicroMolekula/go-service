package controller

import (
	"github.com/MicroMolekula/gpt-service/internal/models"
	"github.com/MicroMolekula/gpt-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChatController struct {
	chatService *service.ChatService
}

func NewChatController(chatService *service.ChatService) *ChatController {
	return &ChatController{
		chatService: chatService,
	}
}

func (c *ChatController) SendMessage(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	type messageRequest struct {
		Text string `json:"text"`
	}
	var message messageRequest
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chatMessage, err := c.chatService.SendMessage(message.Text, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": chatMessage})
}

func (c *ChatController) GetMessages(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	messages, err := c.chatService.GetAllMessages(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}
