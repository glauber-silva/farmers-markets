package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/glauber-silva/farmers-markets/internal/markets"
	"github.com/gorilla/mux"
)

// store pointer to farmers markets service
type Handler struct {
	Router  *mux.Router
	Service *markets.Service
}

// returns a pointer to a Handler
func NewHandler(service *markets.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// setup up all routes for application
func (h *Handler) SetupRoutes() {
	/*
	TODO: Add Search handler
	 */
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/market", h.GetAllMarkets).Methods("GET")
	h.Router.HandleFunc("/api/market", h.PostMarket).Methods("POST")
	h.Router.HandleFunc("/api/market/{id}", h.GetMarket).Methods("GET")
	h.Router.HandleFunc("/api/market/{id}", h.DeleteMarket).Methods("DELETE")
	h.Router.HandleFunc("/api/market/{id}", h.UpdateMarket).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I am alive")
	})
}

// GetMarket Handler - retrieve a market by ID
func (h *Handler) GetMarket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		fmt.Fprint(w, "Unable to parse ID")
	}

	market, err := h.Service.GetMarket(int(i))

	if err != nil {
		fmt.Fprintf(w, "Error retrieving market by ID")
	}

	fmt.Fprintf(w, "%+v", market)
}

// GetAllMarkets Handler - retrieve all markets from the market service
func (h *Handler) GetAllMarkets(w http.ResponseWriter, r *http.Request) {
	markets, err := h.Service.GetAllMarkets()
	if err != nil {
		fmt.Fprintf(w, "Failed retrieving markets")
	}
	fmt.Fprintf(w, "%v", markets)
}

// PostMarket Handler - Add new market
func (h *Handler) PostMarket(w http.ResponseWriter, r *http.Request){
	market, err := h.Service.PostMarket(markets.Market{
		Long: 1234,
		Lat: 54321,
		Bairro: "Centro",
		Logradouro: "Rua Xpto",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to post new market")
	}
	fmt.Fprintf(w, "%+v", market)
}

// UpdateMarket Handler - Update a market by ID
func (h *Handler) UpdateMarket(w http.ResponseWriter, r *http.Request){
	market, err := h.Service.UpdateMarket(1, markets.Market{
		Long: 1234,
		Lat:  4321,
		Bairro: "Centro",
		Logradouro: "Rua Xpto",
	})

	if err != nil {
		fmt.Fprintf(w, "Failed to update market")
	}
	fmt.Fprintf(w, "%+v", market)
}

// DeleteMarket Handler - Delete a market by ID
func (h *Handler) DeleteMarket(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	mktID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Fprint(w, "Unable to parse ID")
	}

	err = h.Service.DeleteMarket(mktID)

	if err != nil {
		fmt.Fprintf(w, "Failed to delete market")
	}

	fmt.Fprintf(w, "Market %+v Deleted successfully", mktID)
}