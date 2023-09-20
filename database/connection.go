package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
	"time"
)

// function to create connection
func CreateConnection() *sql.DB {
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbName := os.Getenv("db_name")

	//dsn := fmt.Sprintf("%v:@tcp(%v:%v)/%v?parseTime=true&loc=Local", dbUser, dbHost, dbPort, dbName)
	dsnSqlServer := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("sqlserver", dsnSqlServer)
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
