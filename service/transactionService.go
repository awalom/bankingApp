package service

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/interfaces"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/model"
	"time"
)

type TransactionService struct {
	interfaces.ITransactionRepo
}

func (ts TransactionService) AddNewTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	err := request.ValidateTransaction(5000.10)
	if err != nil {
		logger.Error("Invalid transaction")
		return nil, errs.ValidationError("Invalid Transaction")
	}
	//convert transaction request to transaction model
	newTransaction := model.Transaction{
		TransactionId:   "",
		AccountId:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	result, resultErr := ts.Save(newTransaction)

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

func GetTransactionService(tr interfaces.ITransactionRepo) TransactionService {
	return TransactionService{tr}
}
