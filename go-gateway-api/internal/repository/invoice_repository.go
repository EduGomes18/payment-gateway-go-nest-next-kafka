package repository

import (
	"database/sql"
	"go-gateway-api/internal/domain"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// Save a new invoice into the database
func (r *InvoiceRepository) Save(invoice *domain.Invoice) error {
	_, err := r.db.Exec(
		"INSERT INTO invoices (id, accountId, amount, status, description, paymentType, cardLastDigits, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		invoice.ID, invoice.AccountID, invoice.Amount, invoice.Status, invoice.Description, invoice.PaymentType, invoice.CardLastDigits, invoice.CreatedAt, invoice.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}


// Find a invoice by id
func (r *InvoiceRepository) FindById(id string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := r.db.QueryRow(`
		SELECT id, accountId, amount, status, description, paymentType, cardLastDigits, createdAt, updatedAt 
		FROM invoices 
		WHERE id = $1
	`, id).Scan(
		&invoice.ID,
		&invoice.AccountID,
		&invoice.Amount,
		&invoice.Status,
		&invoice.Description,
		&invoice.PaymentType,
		&invoice.CardLastDigits,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}

	if err != nil {
		return nil, err
	}

	
	return &invoice, nil
}

// Find all invoices by account id
func (r *InvoiceRepository) FindByAccountId(accountId string) ([]*domain.Invoice, error) {
	rows, err := r.db.Query(`
		SELECT id, accountId, amount, status, description, paymentType, cardLastDigits, createdAt, updatedAt 
		FROM invoices 
		WHERE accountId = $1
	`, accountId)	

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var invoices []*domain.Invoice

	for rows.Next() {
		var invoice domain.Invoice
		err = rows.Scan(&invoice.ID, &invoice.AccountID, &invoice.Amount, &invoice.Status, &invoice.Description, &invoice.PaymentType, &invoice.CardLastDigits, &invoice.CreatedAt, &invoice.UpdatedAt)
		if err != nil {	
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
}	


// UpdateStatus to update the status of a invoice
func (r *InvoiceRepository) UpdateStatus(invoice *domain.Invoice) error {
	rows, err := r.db.Exec(`
		UPDATE invoices 
		SET status = $1 
		WHERE id = $2
	`, invoice.Status, invoice.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrInvoiceNotFound
	}

	return nil
}	

