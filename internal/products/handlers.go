package products

import (
	"log"
	"net/http"

	"github.com/choiexe1/go-ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProduct(w http.ResponseWriter, r *http.Request) {
	// 1. Call service -> ListProducts
	// 2. Return JSON in an HTTP Response

	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, error.Error(err), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}
