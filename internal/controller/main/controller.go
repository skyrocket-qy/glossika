package maincontroller

import (
	"context"

	mainucase "recsvc/internal/usecase/main"
)

type Usecase interface {
	GetRecommendation(ctx context.Context) (*mainucase.GetRecommendationOut, error)
}

type Handler struct {
	Usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}
