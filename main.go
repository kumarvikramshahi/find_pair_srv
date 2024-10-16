package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kumarvikramshahi/streak_assignment/configs"
	"github.com/kumarvikramshahi/streak_assignment/pkg/domain"
	"github.com/kumarvikramshahi/streak_assignment/pkg/repository"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// Load Configuration File
	settingsFileName := os.Args[1]
	configs.LoadConfiguration(
		settingsFileName,
	)

	// init validator
	domain.ValidatorSingletonClient()

	// routes
	route := mux.NewRouter()
	route.HandleFunc(
		"/",
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			err := json.NewEncoder(w).Encode(
				map[string]string{
					"message": "Working...",
				},
			)
			if err != nil {
				return
			}
		},
	)
	repository.NewHttpServer(route)

	// listening
	_ = http.ListenAndServe(":"+configs.ServiceConfiguration.Port, route)
}
