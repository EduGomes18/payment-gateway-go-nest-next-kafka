package service

import (
	"go-gateway-api/internal/domain"
	"go-gateway-api/internal/dto"
)


type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByEmail(account.Email)
	
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrAccountAlreadyExists
	}

	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil

}

func (s *AccountService) UpdateBalance(id string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(id)
	if err != nil {
		return nil, err
	}
	
	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	
	output := dto.FromAccount(account)
	return &output, nil
}


func (s *AccountService) FindByApiKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}
	
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindById(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	
	output := dto.FromAccount(account)
	return &output, nil
}