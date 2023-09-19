package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

// function to create connection
func CreateConnection() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("db_port")
	dbUser := os.Getenv("db_user")
	dbName := os.Getenv("db_name")

	dsn := fmt.Sprintf("%v:@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbUser, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("error cant connect database : " + err.Error())
	}

	fmt.Println("success koneksi ke database")

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	return db
}
