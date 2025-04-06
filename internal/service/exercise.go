package service

import (
	"github.com/MicroMolekula/gpt-service/internal/client"
	"github.com/MicroMolekula/gpt-service/internal/config"
	"github.com/MicroMolekula/gpt-service/internal/dto"
)

type ExerciseService struct {
	exerciseClient *client.ExerciseClient
	cfg            *config.Config
}

func NewExerciseService(cfg *config.Config) *ExerciseService {
	exerciseClient := client.NewExerciseClient(cfg.ExerciseUrl)
	return &ExerciseService{
		exerciseClient: exerciseClient,
	}
}

func (es *ExerciseService) GetExerciseArray(target string, inventory string) ([]string, error) {
	exercises, err := es.exerciseClient.Query(target, inventory)
	if err != nil {
		return nil, err
	}
	arrayEquipments := es.FormatArray(exercises)
	return arrayEquipments, nil
}

func (es *ExerciseService) FormatArray(exercises []dto.ExerciseResponse) []string {
	result := make([]string, len(exercises))
	for i, exercise := range exercises {
		result[i] = exercise.Name
	}
	return result
}
