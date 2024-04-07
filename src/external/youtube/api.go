package youtube

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	httpClient "youtuber/pkg/http"
)

type APIResponder interface {
	MaskCustomerNumber(ctx context.Context, req SearchResourceList) ([]ResourceListResponse, error)
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

func (c *APIClient) MaskCustomerNumber(ctx context.Context, numberMaskingRequest SearchResourceList) ([]ResourceListResponse, error) {

	response := []ResourceListResponse{}

	query := url.Values{}
	query.Add("order", numberMaskingRequest.OrderBy)
	query.Add("publishedAfter", numberMaskingRequest.PublishedAfter)
	query.Add("q", numberMaskingRequest.Query)
	query.Add("type", numberMaskingRequest.ResourceType)
	query.Add("part", "snippet")

	req := httpClient.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/youtube/v3/docs/search/list", c.config.Host), nil, &response).
		WithQuery(query)

	err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
