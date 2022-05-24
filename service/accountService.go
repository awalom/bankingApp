package service

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/repo"
)

// AccountService  The implementation is a struct
type AccountService struct {
	Repo repo.AccountRepo
}

// NewAccount  Receiver function
func (s AccountService) NewAccount(a dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {

	valErr := a.ValidateRequest()
	if valErr != nil {
		logger.Error("Account is not valid")
		return nil, valErr
	}
	account := a.AccountRequestToAccount()

	newAccount, err := s.Repo.Save(account)

	if err != nil {
		logger.Error("Server Error, could not save Account" + a.CustomerId)
		return nil, errs.NewNotFoundError(err.Message)
	}

	logger.InitLogger("Added new account to customer: " + account.CustomerId)
	response := dto.AccountResponse{TransactionId: newAccount.AccountId}

	return &response, nil
}

// GetAccountService  Helper Function
func GetAccountService(repo repo.AccountRepo) AccountService {
	return AccountService{repo}
}
