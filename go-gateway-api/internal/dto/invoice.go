package dto

import (
	"go-gateway-api/internal/domain"
	"time"
)

const (
	StatusPending = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	ApiKey string
	Amount float64 `json:"amount"`
	Description string `json:"description"`
	PaymentType string `json:"paymentType"`
	CardNumber string `json:"cardNumber"`
	CardHolderName string `json:"cardHolderName"`
	CVV string `json:"cvv"`
	ExpirationMonth int `json:"expirationMonth"`
	ExpirationYear int `json:"expirationYear"`
}

type InvoiceOutput struct {
	ID string `json:"id"`
	AccountID string `json:"accountId"`
	Amount float64 `json:"amount"`
	Description string `json:"description"`
	PaymentType string `json:"paymentType"`
	CardLastDigits string `json:"cardLastDigits"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToInvoice(input CreateInvoiceInput, accountId string) (*domain.Invoice, error) {
	card := domain.CreditCard{	
		Number: input.CardNumber,
		HolderName: input.CardHolderName,
		CVV: input.CVV,
		ExpirationMonth: input.ExpirationMonth,
		ExpirationYear: input.ExpirationYear,
	}
	
	
	return domain.NewInvoice(
		accountId,
		input.Amount,
		input.Description,
		input.PaymentType,
		card,
	)
}

func FromInvoice(invoice *domain.Invoice) *InvoiceOutput {
	return &InvoiceOutput{
		ID: invoice.ID,
		AccountID: invoice.AccountID,
		Amount: invoice.Amount,
		Description: invoice.Description,
		PaymentType: invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		Status: string(invoice.Status),
		CreatedAt: invoice.CreatedAt,
		UpdatedAt: invoice.UpdatedAt,
	}
}


