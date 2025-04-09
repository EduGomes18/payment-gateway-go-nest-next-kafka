package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

const MAX_AMOUNT = 10000
const APPROVED_PROBABILITY = 0.7
type Invoice struct {
	ID string
	AccountID string
	Amount float64
	Status Status
	Description string
	PaymentType string
	CardLastDigits string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreditCard struct {
	Number string
	CVV string
	HolderName string
	ExpirationMonth int
	ExpirationYear int
}

func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	if accountID == "" {
		return nil, ErrInvalidAccountID
	}

	if paymentType == "" {
		return nil, ErrInvalidPaymentType
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID: uuid.New().String(),
		AccountID: accountID,
		Amount: amount,
		Status: StatusPending,
		Description: description,
		PaymentType: paymentType,
		CardLastDigits: lastDigits,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > MAX_AMOUNT {
		return nil
	}

	// add random status to simulate the payment gateway

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))

	var newStatus Status

	if randomSource.Float64() <= APPROVED_PROBABILITY {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()

	return nil
}