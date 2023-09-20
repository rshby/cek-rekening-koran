package coresystem

import (
	"database/sql"
)

type TabelCore struct {
	TrDat7         string         `json:"tr_dat_7,omitempty"`
	TrDat6         string         `json:"tr_dat_6,omitempty"`
	TrTime         string         `json:"tr_time,omitempty"`
	TrDorc         string         `json:"tr_dorc,omitempty"`
	Amount         float64        `json:"amount,omitempty"`
	Bal            string         `json:"bal,omitempty"`
	TrStat         string         `json:"tr_stat,omitempty"`
	TrRemk         string         `json:"tr_remk,omitempty"`
	TrRefn         string         `json:"tr_refn,omitempty"`
	EcIndicator    string         `json:"ec_indicator,omitempty"`
	TrRefnEc       string         `json:"tr_refn_ec,omitempty"`
	TellerId       string         `json:"teller_id,omitempty"`
	AtmId          sql.NullString `json:"atm_id,omitempty"`
	AtmJournalSeq  sql.NullString `json:"atm_journal_seq,omitempty"`
	JournalSeq     string         `json:"journal_seq,omitempty"`
	SourceOfBranch string         `json:"source_of_branch,omitempty"`
}
