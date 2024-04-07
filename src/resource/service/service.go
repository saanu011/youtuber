package service

import (
	"context"
	"youtuber/src/external/youtube"
	"youtuber/src/resource/domain"
)

type Service interface {
	GetResourcesList(context.Context, domain.ResourceQuery) ([]domain.Resource, error)
}

type service struct {
	client youtube.APIResponder
}

func NewService(client youtube.APIResponder) Service {
	return &service{client: client}
}

func (s service) GetResourcesList(ctx context.Context, queryRequest domain.ResourceQuery) ([]domain.Resource, error) {
	req := youtube.SearchResourceList{
		PublishedAfter: queryRequest.PublishedAfter,
		Query:          queryRequest.Query,
		OrderBy:        queryRequest.OrderBy,
		ResourceType:   queryRequest.ResourceType,
	}
	res, err := s.client.MaskCustomerNumber(ctx, req)
	if err != nil {
		return []domain.Resource{}, err
	}

	return mapResourcesResponse(res), nil
}

func mapResourcesResponse(res []youtube.ResourceListResponse) []domain.Resource {
	return []domain.Resource{}
}
