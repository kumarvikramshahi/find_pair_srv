package ports

import (
	"github.com/kumarvikramshahi/streak_assignment/pkg/repository/internal"
)

type IRepoCommandService interface {
	Find(
		internal.FindPairRequest,
	) (*internal.FindPairResponse, error)
}
