package sqlClient

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/model"
	"log"
	"time"
)

type SqlClient struct {
	client *sqlx.DB
}

func (d SqlClient) Query() ([]model.Customer, *errs.AppError) {
	var err error
	customers := make([]model.Customer, 0)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	err = d.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil

}

func (d SqlClient) QueryRow(id string) (*model.Customer, *errs.AppError) {

	var c model.Customer

	findOne := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id =?"

	err := d.client.Get(&c, findOne, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer was not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil

}

func GetSqlClient(userName string, password string, dbAddr string, port string, dbname string) SqlClient {
	client, err := sqlx.Open("mysql", getConnectionUrl(userName, password, dbAddr, port, dbname))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return SqlClient{client}
}

func getConnectionUrl(userName string, password string, dbAddr string, port string, dbname string) string {
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, port, dbAddr, dbname)
	return dbConnectionString
}
