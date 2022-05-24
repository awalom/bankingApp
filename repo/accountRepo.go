package repo

import (
	"github.com/jmoiron/sqlx"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/model"
	"strconv"
)

type AccountRepo struct {
	client *sqlx.DB
}

func (accountRepo AccountRepo) Save(a model.Account) (*model.Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	result, err := accountRepo.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while adding new Account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error from database")
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id" + err.Error())
		return nil, errs.NewNotFoundError("Error getting last inserted id")
	}
	a.AccountId = strconv.FormatInt(lastInsertedId, 10)
	return &a, nil
}

func (accountRepo AccountRepo) GetAccount(id string) (*model.Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account model.Account
	err := accountRepo.client.Get(&account, sqlGetAccount, id)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func GetAccountRepo(sqlClient *sqlx.DB) AccountRepo {
	return AccountRepo{sqlClient}
}
