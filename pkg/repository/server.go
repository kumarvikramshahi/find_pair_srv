package repository

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kumarvikramshahi/streak_assignment/pkg/repository/internal/ports"
)

func NewHttpServer(
	router *mux.Router,
) {
	clientPort := ports.NewClientPort()

	router.HandleFunc("/find-pairs", clientPort.FindPair).Methods(http.MethodPost)
}
