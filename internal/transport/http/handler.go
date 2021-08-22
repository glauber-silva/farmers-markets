package http

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/glauber-silva/farmers-markets/internal/markets"
	"github.com/gorilla/mux"
)

// Handler - store pointer to farmers markets service
type Handler struct {
	Router  *mux.Router
	Service *markets.Service
}

// Response - an struct to store responses from API
type Response struct {
	Message string
	Error string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *markets.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path": r.URL.Path,
			}).Info("Handled request")
		next.ServeHTTP(w, r)
	})
}

// SetupRoutes - setup up all routes for application
func (h *Handler) SetupRoutes() {
	/*
		TODO: Add Search handler
	*/
	log.Info("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)
	h.Router.HandleFunc("/api/market", h.GetAllMarkets).Methods("GET")
	h.Router.HandleFunc("/api/market", h.PostMarket).Methods("POST")
	h.Router.HandleFunc("/api/market/{id}", h.GetMarket).Methods("GET")
	h.Router.HandleFunc("/api/market/{id}", h.DeleteMarket).Methods("DELETE")
	h.Router.HandleFunc("/api/market/{id}", h.UpdateMarket).Methods("PUT")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "Alive"}); err != nil {
			panic(err)
		}
	})
}

// GetMarket Handler - retrieve a market by ID
func (h *Handler) GetMarket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		sendErrorResponse(w, "Unable to parse ID", err)
		return
	}

	market, err := h.Service.GetMarket(uint(i))

	if err != nil {
		sendErrorResponse(w, "Error retrieving market by ID", err)
		return
	}

	if err := sendOkResponse(w, market); err != nil {
		panic(err)
	}
}

// GetAllMarkets Handler - retrieve all markets from the market service
func (h *Handler) GetAllMarkets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8" )
	w.WriteHeader(http.StatusOK)

	markets, err := h.Service.GetAllMarkets()
	if err != nil {
		sendErrorResponse(w, "Failed retrieving markets", err)
		return
	}
	if err := json.NewEncoder(w).Encode(markets); err != nil {
		panic(err)
	}
}

// PostMarket Handler - Add new market
func (h *Handler) PostMarket(w http.ResponseWriter, r *http.Request) {
	var market markets.Market
	if err := json.NewDecoder(r.Body).Decode(&market); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	market, err := h.Service.PostMarket(market)
	if err != nil {
		sendErrorResponse(w, "Failed to post new market", err)
		return
	}
	if err := sendOkResponse(w, market); err!= nil {
		panic(err)
	}
}

// UpdateMarket Handler - Update a market by ID
func (h *Handler) UpdateMarket(w http.ResponseWriter, r *http.Request) {
	var market markets.Market
	if err := json.NewDecoder(r.Body).Decode(&market); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	mktID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint from ID")
	}

	market, err = h.Service.UpdateMarket(uint(mktID), market )

	if err != nil {
		sendErrorResponse(w, "Failed to update market", err)
		return
	}

	if err := sendOkResponse(w, market); err != nil {
		panic(err)
	}
}

// DeleteMarket Handler - Delete a market by ID
func (h *Handler) DeleteMarket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	mktID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Unable to parse ID", err)
		return
	}

	err = h.Service.DeleteMarket(uint(mktID))

	if err != nil {
		sendErrorResponse(w, "Failed to delete market", err)
		return
	}

	if err = sendOkResponse(w, Response{Message: "Market deleted successfully"}); err != nil {
		panic(err)
	}

}
func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8" )
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8" )
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(r)
}