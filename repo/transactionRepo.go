package repo

import (
	"github.com/jmoiron/sqlx"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/model"
	"strconv"
)

type TransactionRepo struct {
	client      *sqlx.DB
	accountRepo AccountRepo
}

//func (d TransactionRepo) Save(tr model.Transaction) (*model.Transaction, *errs.AppError) {
//	sqlTransaction := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?,?,?,?)"
//	result, err := d.client.Exec(sqlTransaction, tr.AccountId, tr.Amount, tr.TransactionType, tr.TransactionDate)
//	if err != nil {
//		logger.Error("Error while saving Transaction" + tr.TransactionId)
//		return nil, errs.NewUnexpectedError("Error while saving Transaction" + err.Error())
//	}
//	lastInsertedId, err := result.LastInsertId()
//
//	if err != nil {
//		logger.Error("Error while getting last tr Id" + tr.TransactionId)
//		return nil, errs.NewUnexpectedError("Error while getting last tr Id" + err.Error())
//	}
//
//	tr.TransactionId = strconv.FormatInt(lastInsertedId, 10)
//
//	return &tr, nil
//}

func (d TransactionRepo) Save(tr model.Transaction) (*model.Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	sql := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	result, _ := tx.Exec(sql, tr.AccountId, tr.Amount, tr.TransactionType, tr.TransactionDate)

	// updating account balance
	if tr.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, tr.Amount, tr.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, tr.Amount, tr.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.accountRepo.GetAccount(tr.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	tr.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	tr.Amount = account.Amount
	return &tr, nil
}

func GetTransactionRepo(client *sqlx.DB, accountRepo AccountRepo) TransactionRepo {
	return TransactionRepo{client: client, accountRepo: accountRepo}

}
