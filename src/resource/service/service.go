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
	res, err := s.client.Search(ctx, req)
	if err != nil {
		return []domain.Resource{}, err
	}

	return mapResourcesResponse(res), nil
}

func mapResourcesResponse(res *youtube.ResourceListResponse) []domain.Resource {
	var list []domain.Resource
	for _, e := range res.Data {
		id := domain.ID{
			Kind:    e.ID.Kind,
			VideoID: e.ID.VideoID,
		}

		list = append(list, domain.Resource{
			ID:          id,
			Title:       e.Title,
			Description: e.Description,
			PublishedAt: e.PublishedAt,
			ChannelId:   e.ChannelID,
		})
	}

	return list
}
