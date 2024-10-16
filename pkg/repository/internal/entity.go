package internal

import (
	"encoding/json"
	"io"

	"github.com/kumarvikramshahi/streak_assignment/pkg/domain"
)

type FindPairResponse struct {
	Pairs [][]int `json:"solutions" validate:"required"`
}

type FindPairRequest struct {
	InputArr []int `json:"numbers" validate:"required"`
	Target   int   `json:"target" validate:"required"`
}

func (entity *FindPairRequest) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return err
	}

	err = domain.Validator.Struct(entity)
	if err != nil {
		return err
	}
	return nil
}
