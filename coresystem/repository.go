package coresystem

import (
	"context"
	"database/sql"
	"errors"
)

type CoreSystemRepository struct {
	Db *sql.DB
}

func NewACoreSystemRepository(db *sql.DB) *CoreSystemRepository {
	return &CoreSystemRepository{
		Db: db,
	}
}

// method get data by nomor rekening
func (c *CoreSystemRepository) GetByRekening(ctx context.Context, tx *sql.Tx, accountNumber string) ([]TabelCore, error) {
	// create statement
	queryStatement, err := tx.PrepareContext(ctx, "SELECT DISTINCT A.TRDAT7, A.TRDAT6, A.TRTIME, A.TRDORC, CASE WHEN A.TELLERID IN ('DDCYCLING', 'SVCCHGPAY', 'TAXCHARGE') THEN ROUND(CONVERT(DECIMAL(20,7), A.amount), 2) ELSE A.AMOUNT END AS AMOUNT, '0' BAL, A.TRSTAT, A.TRREMK, A.TRREFN, A.ECINDICATOR, A.TRREFNEC, A.TELLERID, B.ATMID, B.ATMJOURNALSEQ, A.JOURNALSEQ, A.SOURCEOFBRANCH FROM ABCS_T_ABCSHIST A LEFT JOIN ABCS_T_ATMTRN_BAK B ON A.TRREFN = B.TRREFN WHERE A.accountnumber='045202000114806' ORDER BY A.TRDAT7 ASC, A.TRTIME ASC")
	if err != nil {
		return nil, err
	}

	rows, err := queryStatement.QueryContext(ctx)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var results []TabelCore
	for rows.Next() {
		var coreSystem TabelCore
		if err := rows.Scan(&coreSystem.TrDat7, &coreSystem.TrDat6, &coreSystem.TrTime, &coreSystem.TrDorc, &coreSystem.Amount, &coreSystem.Bal, &coreSystem.TrStat, &coreSystem.TrRemk, &coreSystem.TrRefn, &coreSystem.EcIndicator, &coreSystem.TrRefnEc, &coreSystem.TellerId, &coreSystem.AtmId, &coreSystem.AtmJournalSeq, &coreSystem.JournalSeq, &coreSystem.SourceOfBranch); err != nil {
			return nil, err
		}
		results = append(results, coreSystem)
	}

	if len(results) == 0 {
		return nil, errors.New("record not found")
	}

	return results, nil
}
