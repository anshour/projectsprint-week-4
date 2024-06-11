package db

import (
	"context"
	"fmt"
	"projectsprintw4/src/config"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func Init() *sqlx.DB {
	var isInitialized bool

	if isInitialized {
		return nil
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		config.DB_PARAMS,
	)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	db, err := sqlx.ConnectContext(ctx, "pgx", connStr)
	if err != nil {
		fmt.Println("Error creating database connection: ", err.Error())
		panic(err.Error())
	}

	fmt.Println("Database connected!")

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(90)
	db.SetConnMaxIdleTime(time.Second * 5)
	db.SetConnMaxLifetime(time.Hour)

	isInitialized = true

	return db
}
