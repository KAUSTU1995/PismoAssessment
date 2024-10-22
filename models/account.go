package models

// Account represents the account model for the PismoAssessment API.
// @Description Account model for managing account information.
type Account struct {
	// AccountID is the unique identifier for the account.
	// @example 1
	AccountID int `json:"account_id"`

	// DocumentNumber is the document number associated with the account.
	// @example "12345678900"
	// @validate required,len=11
	DocumentNumber string `json:"document_number" validate:"required,len=11"`
}
