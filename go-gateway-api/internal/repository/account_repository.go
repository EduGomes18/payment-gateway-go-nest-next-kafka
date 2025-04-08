package repository

import (
	"database/sql"
	"time"

	"go-gateway-api/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	smt, err := r.db.Prepare(`
	INSERT INTO accounts (id, name, email, apiKey, balance, createdAt, updatedAt) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindByEmail(email string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time


	err := r.db.QueryRow(`
	SELECT ID, name, email, apiKey, balance, createdAt, updatedAt 
	FROM accounts WHERE email = $1
	`, email).Scan(
		&account.ID,
		&account.Name,		
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)
	

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt		

	return &account, nil
}

func (r *AccountRepository) FindByApiKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time


	err := r.db.QueryRow(`
	SELECT ID, name, email, apiKey, balance, createdAt, updatedAt 
	FROM accounts WHERE apiKey = $1
	`, apiKey).Scan(
		&account.ID,
		&account.Name,		
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)
	

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt		

	return &account, nil
}

func (r *AccountRepository) FindById(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time


	err := r.db.QueryRow(`
	SELECT ID, name, email, apiKey, balance, createdAt, updatedAt 
	FROM accounts 
	WHERE id = $1
	`, id).Scan(
		&account.ID,
		&account.Name,		
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)
	

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt		

	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	var currentBalance float64
	
	err = tx.QueryRow(`
	SELECT balance FROM accounts WHERE id = $1 FOR UPDATE
	`, account.ID).Scan(&currentBalance)
	
	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}

	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	UPDATE accounts SET balance = $1, updatedAt = $2 WHERE id = $3
	`, account.Balance, time.Now(), account.ID)

	if err != nil {
		return err
	}

	return tx.Commit()

}
