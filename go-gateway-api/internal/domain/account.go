package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mutex     sync.RWMutex // Mutex to protect balance from multi write operations
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateApiKey() string {
	apiKey := make([]byte, 16)
	rand.Read(apiKey)
	return hex.EncodeToString(apiKey)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		APIKey: generateApiKey(),
		Balance: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) AddBalance(amount float64) {
	a.mutex.Lock() // Lock the mutex to prevent race conditions
	defer a.mutex.Unlock() // Unlock the mutex after the function returns -- this is a good practice to avoid deadlocks, 
	// also its runs at the end of the function after the function is executed	
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) GetBalance() float64 {
	a.mutex.RLock() // Lock the mutex to prevent race conditions
	defer a.mutex.RUnlock() // Unlock the mutex after the function returns
	return a.Balance
}
