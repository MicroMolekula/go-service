package controller

import (
	"github.com/MicroMolekula/gpt-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GPTController struct {
	gptService *service.GptService
}

func NewGPTController(gptService *service.GptService) *GPTController {
	return &GPTController{
		gptService: gptService,
	}
}

func (c *GPTController) GetAnswer(ctx *gin.Context) {
	type gptRequest struct {
		System string `json:"system"`
		User   string `json:"user"`
	}
	var req gptRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.gptService.Query(req.System, req.User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
