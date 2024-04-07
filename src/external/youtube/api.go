package youtube

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	httpClient "youtuber/pkg/http"
)

type APIResponder interface {
	Search(ctx context.Context, req SearchResourceList) (*ResourceListResponse, error)
}

type APIClient struct {
	config     Config
	httpClient *httpClient.Client
}

func NewAPIClient(cfg Config) *APIClient {
	return &APIClient{
		config:     cfg,
		httpClient: httpClient.New(cfg.APICommandName),
	}
}

func (c *APIClient) Search(ctx context.Context, request SearchResourceList) (*ResourceListResponse, error) {

	response := &ResourceListResponse{}

	query := url.Values{}
	query.Add("order", request.OrderBy)
	query.Add("publishedAfter", request.PublishedAfter)
	query.Add("q", request.Query)
	query.Add("type", request.ResourceType)
	query.Add("part", "snippet")
	query.Add("key", c.config.AuthKey)

	req := httpClient.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/youtube/v3/search", c.config.Host), nil, &response).
		WithQuery(query)

	err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
