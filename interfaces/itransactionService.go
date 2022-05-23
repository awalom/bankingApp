package interfaces

import (
	"gitlab/awalom/banking/dto"
	"gitlab/awalom/banking/errs"
)

type ITransactionService interface {
	AddNewTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}
