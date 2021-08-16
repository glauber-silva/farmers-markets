package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Store pointer to farmers markets service
type Handler struct {
	Router *mux.Router
}

// returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// setup up all routes for application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I am alive")
	})
}
