package client

import (
	"bytes"
	"encoding/json"
	"github.com/MicroMolekula/gpt-service/internal/dto"
	"io"
	"net/http"
)

type ExerciseClient struct {
	httpClient *http.Client
	baseUrl    string
}

func NewExerciseClient(baseUrl string) *ExerciseClient {
	httpClient := &http.Client{}
	return &ExerciseClient{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

func (client *ExerciseClient) Query(target string, inventory string) ([]dto.ExerciseResponse, error) {
	var buf = &bytes.Buffer{}
	body := map[string]interface{}{
		"types":       []string{target},
		"inventories": []string{inventory},
	}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.baseUrl, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = CheckResponse(res); err != nil {
		return nil, err
	}

	var result []dto.ExerciseResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}
