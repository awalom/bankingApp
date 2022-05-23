package repo

import (
	"github.com/jmoiron/sqlx"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/model"
	"strconv"
)

type TransactionRepo struct {
	*sqlx.DB
}

func (t TransactionRepo) Save(tr model.Transaction) (*model.Transaction, *errs.AppError) {
	sqlTransaction := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?,?,?,?)"
	result, err := t.Exec(sqlTransaction, tr.AccountId, tr.Amount, tr.TransactionType, tr.TransactionDate)
	if err != nil {
		logger.Error("Error while saving Transaction" + tr.TransactionId)
		return nil, errs.NewUnexpectedError("Error while saving Transaction" + err.Error())
	}
	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last tr Id" + tr.TransactionId)
		return nil, errs.NewUnexpectedError("Error while getting last tr Id" + err.Error())
	}

	tr.TransactionId = strconv.FormatInt(lastInsertedId, 10)

	return &tr, nil
}

func GetTransactionRepo(client *sqlx.DB) TransactionRepo {
	return TransactionRepo{client}
}
