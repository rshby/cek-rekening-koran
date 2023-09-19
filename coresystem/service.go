package coresystem

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

type AgroCoreService struct {
	Db   *sql.DB
	Repo *AgroCoreRepository
}

func NewAgroCoreService(db *sql.DB, repo *AgroCoreRepository) *AgroCoreService {
	return &AgroCoreService{
		Db:   db,
		Repo: repo,
	}
}

// method process
func (s *AgroCoreService) Process(ctx context.Context, accountNumber string) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// get data
	rows, err := s.Repo.GetByAccountNumber(ctx, tx, accountNumber)
	if err != nil {
		tx.Rollback()
		return err
	}

	// create to excel file
	file := excelize.NewFile()
	defer file.Close()

	// create new sheet
	_, err = file.NewSheet("sheet1")
	if err != nil {
		tx.Rollback()
		return err
	}

	file.SetCellValue("sheet1", "A1", "No")
	file.SetCellValue("sheet1", "B1", "account_number")
	file.SetCellValue("sheet1", "C1", "account_name")
	file.SetCellValue("sheet1", "D1", "dorc")
	file.SetCellValue("sheet1", "E1", "remark")
	file.SetCellValue("sheet1", "F1", "amount")
	file.SetCellValue("sheet1", "G1", "saldo_saat_ini")

	// isi file
	counter := 1
	var saldo float64
	for idx, item := range rows {
		if strings.ToLower(item.Dorc) == "c" {
			saldo += item.Amount
		} else {
			saldo -= item.Amount
		}

		file.SetCellValue("sheet1", fmt.Sprintf("A%v", idx+2), counter)
		file.SetCellValue("sheet1", fmt.Sprintf("B%v", idx+2), item.AccountNumber)
		file.SetCellValue("sheet1", fmt.Sprintf("C%v", idx+2), item.AccountName)
		file.SetCellValue("sheet1", fmt.Sprintf("D%v", idx+2), item.Dorc)
		file.SetCellValue("sheet1", fmt.Sprintf("E%v", idx+2), item.Remark)
		file.SetCellValue("sheet1", fmt.Sprintf("F%v", idx+2), item.Amount)
		file.SetCellValue("sheet1", fmt.Sprintf("G%v", idx+2), saldo)
		counter++
	}

	// save file
	err = file.SaveAs("./output/contoh2.xlsx")
	if err != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()
	return nil
}
