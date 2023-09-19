package main

import (
	"cek-rekening-koran/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf(fmt.Sprintf("error cant load env: %v", err.Error()))
	} else {
		fmt.Println("cek env db_host :", os.Getenv("DB_HOST"))
	}
}

func main() {
	fmt.Println("== app run ==")

	db := database.CreateConnection()
	defer db.Close()
}
