package domain


type AccountRepository interface {
	Save(account *Account) error
	FindByEmail(email string) (*Account, error)
	FindByApiKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	UpdateBalance(account *Account) error
}

