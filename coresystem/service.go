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
	Repo *CoreSystemRepository
}

func NewAgroCoreService(db *sql.DB, repo *CoreSystemRepository) *AgroCoreService {
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
	rows, err := s.Repo.GetByRekening(ctx, tx, accountNumber)
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
	file.SetCellValue("sheet1", "B1", "TrDat7")
	file.SetCellValue("sheet1", "C1", "TrDat6")
	file.SetCellValue("sheet1", "D1", "TrTime")
	file.SetCellValue("sheet1", "E1", "TrDorc")
	file.SetCellValue("sheet1", "F1", "Amount")
	file.SetCellValue("sheet1", "G1", "saldo_saat_ini")
	file.SetCellValue("sheet1", "H1", "TrRemk")
	file.SetCellValue("sheet1", "I1", "TrRefn")
	file.SetCellValue("sheet1", "J1", "EcIndicator")
	file.SetCellValue("sheet1", "K1", "TrRefnEc")
	file.SetCellValue("sheet1", "L1", "TellerId")
	file.SetCellValue("sheet1", "M1", "AtmId")
	file.SetCellValue("sheet1", "N1", "AtmJournalSeq")
	file.SetCellValue("sheet1", "O1", "JournalSeq")
	file.SetCellValue("sheet1", "P1", "SourceOfBranch")
	file.SetCellValue("sheet1", "Q1", "TrStat")

	// isi file
	counter := 1
	var saldo float64
	for idx, item := range rows {
		if strings.ToLower(item.TrDorc) == "c" {
			if strings.Trim(item.EcIndicator, " ") == "1" {
				saldo -= item.Amount
			} else {
				saldo += item.Amount
			}
		} else {
			if strings.Trim(item.EcIndicator, " ") == "1" {
				saldo += item.Amount
			} else {
				saldo -= item.Amount
			}
		}

		file.SetCellValue("sheet1", fmt.Sprintf("A%v", idx+2), counter)
		file.SetCellValue("sheet1", fmt.Sprintf("B%v", idx+2), item.TrDat7)
		file.SetCellValue("sheet1", fmt.Sprintf("C%v", idx+2), item.TrDat6)
		file.SetCellValue("sheet1", fmt.Sprintf("D%v", idx+2), item.TrTime)
		file.SetCellValue("sheet1", fmt.Sprintf("E%v", idx+2), item.TrDorc)
		file.SetCellValue("sheet1", fmt.Sprintf("F%v", idx+2), item.Amount)
		file.SetCellValue("sheet1", fmt.Sprintf("G%v", idx+2), saldo)
		file.SetCellValue("sheet1", fmt.Sprintf("H%v", idx+2), item.TrRemk)
		file.SetCellValue("sheet1", fmt.Sprintf("I%v", idx+2), item.TrRefn)
		file.SetCellValue("sheet1", fmt.Sprintf("J%v", idx+2), item.EcIndicator)
		file.SetCellValue("sheet1", fmt.Sprintf("K%v", idx+2), item.TrRefnEc)
		file.SetCellValue("sheet1", fmt.Sprintf("L%v", idx+2), item.TellerId)
		file.SetCellValue("sheet1", fmt.Sprintf("M%v", idx+2), item.AtmId.String)
		file.SetCellValue("sheet1", fmt.Sprintf("N%v", idx+2), item.AtmJournalSeq.String)
		file.SetCellValue("sheet1", fmt.Sprintf("O%v", idx+2), item.JournalSeq)
		file.SetCellValue("sheet1", fmt.Sprintf("P%v", idx+2), item.SourceOfBranch)
		file.SetCellValue("sheet1", fmt.Sprintf("Q%v", idx+2), item.TrStat)
		counter++
	}

	// save file
	err = file.SaveAs("./output/contoh5.xlsx")
	if err != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()
	return nil
}
