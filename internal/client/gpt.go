package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/MicroMolekula/gpt-service/internal/dto"
	"io"
	"net/http"
	"net/url"
	"time"
)

type GPTClient struct {
	httpClient *http.Client
	baseUrl    *url.URL
}

func NewGPTClient(baseUrl string) *GPTClient {
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	transport := &http.Transport{
		TLSHandshakeTimeout: 2 * time.Minute,
		TLSClientConfig:     tlsConfig,
	}
	httpClient := &http.Client{Transport: transport, Timeout: 2 * time.Minute}
	path, _ := url.Parse(baseUrl)
	return &GPTClient{httpClient: httpClient, baseUrl: path}
}

func (c *GPTClient) NewRequest(
	apiKey string,
	catalogId string,
	temperature float64,
	maxTokens int,
	systemPrompt, userPrompt string,
) (*http.Request, error) {
	body := &dto.GptRequest{
		ModelUri: fmt.Sprintf("gpt://%s/yandexgpt/latest", catalogId),
		CompletionOptions: &dto.CompletionOptions{
			Stream:      false,
			Temperature: temperature,
			MaxTokens:   maxTokens,
			ReasoningOptions: &dto.ReasoningOptions{
				Mode: "DISABLED",
			},
		},
		Messages: []*dto.GptMessage{
			{
				Role: "system",
				Text: systemPrompt,
			},
			{
				Role: "user",
				Text: userPrompt,
			},
		},
	}
	var buf = &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}
	fmt.Println(buf.String())
	req, err := http.NewRequest("POST", c.baseUrl.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Api-Key "+apiKey)

	return req, nil
}

func (c *GPTClient) Do(req *http.Request, v any) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = CheckResponse(resp); err != nil {
		return resp, err
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return fmt.Errorf("%s %s: %s", resp.Request.Method, resp.Request.URL, resp.Status)
}
