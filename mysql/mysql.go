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

func init() {
	go func() {
		for {
			<-time.After(1 * time.Second)
			if db == nil {
				continue
			}

			if err := db.Ping(); err != nil {
				continue
			}

			stat := db.Stats()
			log.Debug("open-conn=", stat.OpenConnections, "\tin-use=", stat.InUse, "\tidle=", stat.Idle, "\twait=", stat.WaitCount)
		}
	}()
}

func open(user, password, dbname string, maxConn, maxIdleConn int, connMaxLife time.Duration) (*sql.DB, error) {
	log.Info("init db with parameters\t", maxConn, maxIdleConn, connMaxLife)
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
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(connMaxLife)
	db.SetConnMaxIdleTime(0)

	return db, nil
}

func Query(q string, args ...any) (*sql.Rows, error) {
	if db == nil {
		return nil, ErrDatabaseIsNotInitialized
	}

	ctx, cancel := context.WithTimeout(context.Background(), getTimeout())
	defer cancel()

	return db.QueryContext(ctx, q, args...)
}

func Execute(q string, args ...any) error {
	if db == nil {
		return ErrDatabaseIsNotInitialized
	}

	ctx, cancel := context.WithTimeout(context.Background(), getTimeout())
	defer cancel()

	_, err := db.ExecContext(ctx, q, args...)

	return err
}

func Init(user, password, dbname string, maxConn, maxIdleConn int, connMaxLife time.Duration) {
	log.Debug("Initial database")
	_, err := open(user, password, dbname, maxConn, maxIdleConn, connMaxLife)
	if err != nil {
		panic(err)
	}
}

func InitWithConfig() {
	Init(Config())
}
