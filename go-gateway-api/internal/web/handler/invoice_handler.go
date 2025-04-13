package handler

import (
	"encoding/json"
	"go-gateway-api/internal/domain"
	"go-gateway-api/internal/dto"
	"go-gateway-api/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)


type InvoiceHandler struct {
	invoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}


// Request with authentication via api key
// Endpoint: /invoices
// Method: POST
// Body: CreateInvoiceInput
// Response: InvoiceOutput
func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.ApiKey = r.Header.Get("X-API-KEY")

	output, err := h.invoiceService.CreateInvoice(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}


// Request with authentication via api key
// Endpoint: /invoices/:id
// Method: GET
// Response: InvoiceOutput
func (h *InvoiceHandler) GetInvoiceById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// log.Printf("GetInvoiceById: %s", id)
	if id == "" {
		http.Error(w, "Invoice ID is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusUnauthorized)
		return
	}

	output, err := h.invoiceService.GetById(id, apiKey)
	if err != nil {
		switch err {
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case domain.ErrUnauthorized:
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}


// Request with authentication via api key
// Endpoint: /invoices
// Method: GET
// list by account api key
// Response: []InvoiceOutput
func (h *InvoiceHandler) ListInvoicesByAccountApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "API key is required", http.StatusUnauthorized)
		return
	}

	output, err := h.invoiceService.ListByAccountApiKey(apiKey)
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

