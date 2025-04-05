package service

import (
	"context"
	"encoding/json"
	"github.com/MicroMolekula/gpt-service/internal/config"
	"github.com/MicroMolekula/gpt-service/internal/dto"
	"github.com/MicroMolekula/gpt-service/internal/models"
	"github.com/MicroMolekula/gpt-service/internal/repository"
	"github.com/MicroMolekula/gpt-service/internal/utils"
	"strconv"
	"strings"
)

type FitnessService struct {
	cfg                *config.Config
	gptService         *GptService
	userPlanRepository *repository.UserPlanRepository
}

func NewFitnessService(gptService *GptService, cfg *config.Config, userPlanRepository *repository.UserPlanRepository) *FitnessService {
	return &FitnessService{gptService: gptService, cfg: cfg, userPlanRepository: userPlanRepository}
}

func (fs *FitnessService) GetPlanByUserId(user *models.User) (*dto.WeekPlan, error) {
	userId := strconv.Itoa(int(user.ID))
	userPlan, err := fs.userPlanRepository.GetByUserID(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	weekPlan := &dto.WeekPlan{
		Plan: userPlan.Plan,
	}
	return weekPlan, nil
}

func (fs *FitnessService) GeneratePlanByUser(user *models.User) (*dto.WeekPlan, error) {
	query := utils.GenerateQueryByUserData(user)
	weekPlan, err := fs.GeneratePlan(query)
	if err != nil {
		return nil, err
	}
	userPlan := dto.UserPlan{
		UserId: strconv.Itoa(int(user.ID)),
		Plan:   weekPlan.Plan,
	}
	if err = fs.userPlanRepository.CreateOrUpdate(context.Background(), userPlan); err != nil {
		return nil, err
	}
	return weekPlan, nil
}

func (fs *FitnessService) GeneratePlan(target string) (*dto.WeekPlan, error) {
	result, err := fs.gptService.Query(
		fs.cfg.Prompts.Plans,
		target,
	)
	if err != nil {
		return nil, err
	}
	var data dto.WeekPlan
	text := formatResult(result.Alternatives[0].GptMessage.Text)
	err = json.Unmarshal([]byte(text), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func formatResult(data string) string {
	result := strings.Trim(data, "```")
	return result
}
