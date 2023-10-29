package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tnuanchuay/honeypot/log"
	"strconv"
	"time"
)

var db *sql.DB

var ErrDatabaseIsNotInitialized = errors.New("the database is not initialized")

func open(user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@/", user, password)
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	err = Execute(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", dbname))
	if err != nil {
		return nil, err
	}

	connStr = fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func Execute(query string, args ...interface{}) error {
	if db == nil {
		return ErrDatabaseIsNotInitialized
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		return err
	}
	defer conn.Close()

	timeout, err := strconv.Atoi(getTimeout())
	if err != nil {
		timeout = 60000
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))

	if len(args) == 0 {
		_, err = conn.ExecContext(ctx, query)
	} else {
		_, err = conn.ExecContext(ctx, query, args)
	}
	if err != nil {
		return err
	}

	return nil
}

func Init(user, password, dbname string) {
	log.Debug("Initial database")
	_, err := open(user, password, dbname)
	if err != nil {
		panic(err)
	}
}

func InitWithDefault() {
	Init(Config())
}
