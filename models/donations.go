package models

type DonationsResponse struct {
	Id           int64  `db:"trans_id"`
	Amount       string `db:"donation_amount_"`
	Reference    string `db:"reference_"`
	Email        string `db:"email_"`
	Date_Donated string `db:"date_donated_"`
}

//omitempty
