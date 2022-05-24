package service

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
)

type TransactionService struct {
	tr interfaces.ITransactionRepo
	ac interfaces.IAccountRepo
}

func (ts TransactionService) AddNewTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	//request validation
	err := request.ValidateTransaction()
	if err != nil {
		return nil, errs.NewUnavailable(err.Message)
	}

	//check if fund exist in the account
	account, err := ts.ac.GetAccount(request.AccountId)
	if err != nil {
		return nil, errs.NewUnavailable(err.Message)
	}
	if account.Amount < request.Amount {
		return nil, errs.NewNotFoundError("Balance is to low")
	}

	//convert transaction request to transaction model
	newTransaction := request.TransactionRequestToTransaction()

	result, resultErr := ts.tr.Save(newTransaction)

	if resultErr != nil {
		logger.Error("Error while saving new transaction" + newTransaction.AccountId + resultErr.Message)
		return nil, errs.NewUnexpectedError("Error while saving new transaction")
	}

	//convert transaction model to transaction response
	newTransactionResponse := dto.TransactionResponse{
		TransactionId: result.TransactionId,
		Balance:       result.Amount,
	}

	return &newTransactionResponse, nil

}

func GetTransactionService(tr interfaces.ITransactionRepo, ar interfaces.IAccountRepo) TransactionService {
	return TransactionService{tr, ar}
}
