package controller

import (
	"github.com/MicroMolekula/gpt-service/internal/models"
	"github.com/MicroMolekula/gpt-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FitnessController struct {
	fitnessService *service.FitnessService
}

func NewFitnessController(fitnessService *service.FitnessService) *FitnessController {
	return &FitnessController{fitnessService: fitnessService}
}

func (f *FitnessController) GetPlanWithComment(ctx *gin.Context) {
	type comment struct {
		Comment string `json:"comment"`
	}
	var req comment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := ctx.MustGet("user").(*models.User)
	res, err := f.fitnessService.GeneratePlanByUser(user, req.Comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (f *FitnessController) GetPlan(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	res, err := f.fitnessService.GeneratePlanByUser(user, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (f *FitnessController) FindPlanByUserId(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	res, err := f.fitnessService.GetPlanByUserId(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
