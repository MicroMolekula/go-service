package service

import (
	"github.com/MicroMolekula/gpt-service/internal/client"
	"github.com/MicroMolekula/gpt-service/internal/config"
	"github.com/MicroMolekula/gpt-service/internal/dto"
)

type GptService struct {
	cfg       *config.Config
	gptClient *client.GPTClient
}

func NewGptService(cfg *config.Config) *GptService {
	return &GptService{
		cfg:       cfg,
		gptClient: client.NewGPTClient(cfg.YandexGPT.URL),
	}
}

func (s *GptService) QueryLite(systemPrompt string, userPrompt string) (*dto.GptResult, error) {
	req, err := s.gptClient.NewRequest(
		s.cfg.YandexGPT.ApiToken,
		s.cfg.YandexGPT.CatalogToken,
		0.4,
		200,
		systemPrompt,
		userPrompt,
		true,
	)
	if err != nil {
		return nil, err
	}
	var res = &dto.GptResponse{}
	_, err = s.gptClient.Do(req, res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *GptService) Query(systemPrompt string, userPrompt string) (*dto.GptResult, error) {
	req, err := s.gptClient.NewRequest(
		s.cfg.YandexGPT.ApiToken,
		s.cfg.YandexGPT.CatalogToken,
		0.5,
		3000,
		systemPrompt,
		userPrompt,
		false,
	)
	if err != nil {
		return nil, err
	}
	var res = &dto.GptResponse{}
	_, err = s.gptClient.Do(req, res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
