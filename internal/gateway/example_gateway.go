package gateway

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/yourusername/go-skeleton/internal/model"
)

type ExampleGateway interface {
	SendToExternalAPI(ctx context.Context, data *model.ExampleRequest) (*model.ExampleResponse, error)
}

type exampleGateway struct {
	httpClient *http.Client
	baseURL    string
}

func NewExampleGateway(baseURL string) ExampleGateway {
	return &exampleGateway{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (g *exampleGateway) SendToExternalAPI(ctx context.Context, data *model.ExampleRequest) (*model.ExampleResponse, error) {
	// Example implementation of calling external HTTP API
	// This is a placeholder - replace with actual implementation

	url := g.baseURL + "/api/example"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response model.ExampleResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	_ = jsonData // Use the jsonData variable to avoid unused variable error

	return &response, nil
}
