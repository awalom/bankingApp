package sqlClient

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gitlab/awalom/banking/errs"
	"gitlab/awalom/banking/logger"
	"gitlab/awalom/banking/model"
	"log"
	"time"
)

type SqlClient struct {
	client *sql.DB
}

func (d SqlClient) Query() ([]model.Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	customers := make([]model.Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {
		logger.Error("Error while scanning customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil

}

func (d SqlClient) QueryRow(id string) (*model.Customer, *errs.AppError) {

	var c model.Customer

	findOne := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id =?"

	row := d.client.QueryRow(findOne, id)

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
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

func GetSqlClient(userName string, password string, dbProtocol string, connectionUrl string, dbname string) SqlClient {
	client, err := sql.Open("mysql", getConnectionUrl(userName, password, dbProtocol, connectionUrl, dbname))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return SqlClient{client}
}

func getConnectionUrl(userName string, password string, dbProtocol string, connectionUrl string, dbname string) string {
	return userName + ":" + password + "@" + dbProtocol + "(" + connectionUrl + ")" + "/" + dbname
}
