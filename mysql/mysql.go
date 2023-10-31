package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
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

//func getTimeout () time.Duration {
//	t, err := strconv.Atoi(getTimeout())
//	if err != nil {
//		return 60 * time.Second
//	}
//
//	timeout := time.Duration(t) * time.Millisecond
//}

func Query(q string, args ...any) (*sql.Rows, error) {
	if db == nil {
		return nil, ErrDatabaseIsNotInitialized
	}

	ctx, cancel := context.WithTimeout(context.Background(), getTimeout())
	defer cancel()

	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	return conn.QueryContext(ctx, q, args...)
}

func Execute(q string, args ...any) error {
	if db == nil {
		return ErrDatabaseIsNotInitialized
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), getTimeout())
	defer cancel()

	_, err = conn.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return ctx.Err()
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
