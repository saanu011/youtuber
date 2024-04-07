package handler

import (
	"net/http"
	"youtuber/src/resource/domain"
	"youtuber/src/resource/service"

	response "youtuber/pkg/http"
)

type Handler struct {
	cfg    Config
	client service.Service
}

func NewHandler(cnf Config, client service.Service) *Handler {
	return &Handler{
		cfg:    cnf,
		client: client,
	}
}

func (h *Handler) GetVideoListByParams(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	orderBy := query.Get("data")
	if len(orderBy) == 0 {
		orderBy = "date"
	}
	resourceType := query.Get("resource_type")
	if len(orderBy) == 0 {
		resourceType = "video"
	}

	res, err := h.client.GetResourcesList(r.Context(), domain.ResourceQuery{
		PublishedAfter: query.Get("publishedAfter"),
		Query:          query.Get("query"),
		OrderBy:        orderBy,
		ResourceType:   resourceType,
	})
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusOK, res)

}
