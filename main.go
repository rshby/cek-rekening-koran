package main

import (
	"cek-rekening-koran/coresystem"
	"cek-rekening-koran/database"
	"context"
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

	// create repository layer
	repo := coresystem.NewACoreSystemRepository(db)

	/*
			tx, err := db.Begin()
			if err != nil {
				log.Fatalf(err.Error())
			}

			defer func() {
				if err := recover(); err != nil {
					tx.Rollback()
				}
			}()



		results, err := repo.GetByRekening(context.Background(), tx, "045202000114806")
		fmt.Println(results)

	*/

	//create service layer
	service := coresystem.NewAgroCoreService(db, repo)

	err := service.Process(context.Background(), "301137167")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("sukes create file")
}
