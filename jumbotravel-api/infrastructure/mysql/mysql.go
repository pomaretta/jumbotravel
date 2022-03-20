package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/domain"
	"github.com/pomaretta/jumbotravel/jumbotravel-api/infrastructure/builders"
)

type MySQL struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string

	con *sql.DB
}

func (db *MySQL) Connect() (err error) {

	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DatabaseName,
	)

	db.con, err = sql.Open("mysql", connStr)

	db.con.SetConnMaxLifetime(time.Minute * 3)
	db.con.SetMaxOpenConns(10)
	db.con.SetMaxIdleConns(10)

	return
}

func (db *MySQL) Disconnect() error {
	return db.con.Close()
}

func (db *MySQL) Fetch(ent domain.Entity, qb builders.Builder) ([]interface{}, error) {
	query, args, err := qb.BuildQuery()
	if err != nil {
		return nil, err
	}

	return db.getSlice(ent, query, args...)
}

func (db *MySQL) getSlice(ent domain.Entity, q string, args ...interface{}) (slice []interface{}, err error) {
	rows, err := db.con.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ent.New()
		destFields := ent.GetDestFields()
		if err = rows.Scan(destFields...); err != nil {
			return nil, err
		}
		slice = append(slice, ent.Val())
	}

	return
}

func (db *MySQL) getIntValue(q string, args ...interface{}) (value int, err error) {
	rows, err := db.con.Query(q, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		return
	}
	if err = rows.Scan(&value); err != nil {
		return
	}

	return
}

func (db *MySQL) FetchCount(qb builders.Builder) (int, error) {
	query, args, err := qb.BuildQuery()
	if err != nil {
		return -1, err
	}

	return db.getIntValue(query, args...)
}

func (db *MySQL) Put(qb builders.Builder) (int64, error) {
	query, args, err := qb.BuildQuery()
	if err != nil {
		return -1, err
	}
	res, err := db.con.Exec(query, args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (db *MySQL) Update(qb builders.Builder) (int64, error) {
	query, args, err := qb.BuildQuery()
	if err != nil {
		return -1, err
	}
	res, err := db.con.Exec(query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}
