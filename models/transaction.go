package models

// Transaction represents the transaction model for the PismoAssessment API.
// @Description Transaction model for managing transaction information.
type Transaction struct {
	// TransactionID is the unique identifier for the transaction.
	// @example 1
	TransactionID int `json:"transaction_id"`

	// AccountID is the identifier for the associated account.
	// @example 1
	AccountID int `json:"account_id" validate:"required,gt=0"`

	// OperationTypeID is the identifier for the type of operation performed.
	// @example 2
	OperationTypeID int `json:"operation_type_id" validate:"required,gt=0,operationtype"`

	// Amount is the amount involved in the transaction.
	// @example 100.50
	Amount float64 `json:"amount" validate:"required,gt=0"`

	// EventDate is the date of the transaction event.
	// @example "2024-10-22T15:04:05Z"
	EventDate string `json:"event_date"`
}
