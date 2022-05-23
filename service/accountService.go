package service

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/model"
	"time"
)

// AccountService  The implementation is a struct
type AccountService struct {
	Repo interfaces.IAccountRepo
}

// NewAccount  Receiver function
func (s AccountService) NewAccount(a dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {

	valErr := a.ValidateRequest()
	if valErr != nil {
		logger.Error("Account is not valid")
		return nil, valErr
	}
	account := model.Account{
		AccountId:   "",
		CustomerId:  a.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      "1",
	}

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
func GetAccountService(repository interfaces.IAccountRepo) AccountService {
	return AccountService{repository}
}
