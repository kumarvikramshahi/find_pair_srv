package ports

import (
	"net/http"

	"github.com/kumarvikramshahi/streak_assignment/pkg/domain"
	"github.com/kumarvikramshahi/streak_assignment/pkg/repository/internal"
	"github.com/kumarvikramshahi/streak_assignment/pkg/repository/internal/services"
)

type ClientPort struct {
	IRepoCommandService IRepoCommandService
}

func NewClientPort() *ClientPort {
	port := &ClientPort{
		IRepoCommandService: services.NewFindPairService(),
	}
	return port
}

func (port *ClientPort) FindPair(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var requestDTO internal.FindPairRequest
	if err := requestDTO.Populate(request.Body); err != nil {
		errEntity := &domain.Entity{
			Message:        err.Error(),
			Code:           domain.InvalidRequestCode,
			HttpStatusCode: http.StatusBadRequest,
		}
		domain.SendHttpError(responseWriter, errEntity)
		return
	}
	resultPairs, err := port.IRepoCommandService.Find(requestDTO)
	if err != nil {
		errEntity := &domain.Entity{
			Message:        err.Error(),
			Code:           domain.ServerErrorCode,
			HttpStatusCode: http.StatusInternalServerError,
		}
		domain.SendHttpError(responseWriter, errEntity)
		return
	}
	domain.SendHttpResponse(responseWriter, http.StatusOK, resultPairs)
}
