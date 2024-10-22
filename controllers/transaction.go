package controllers

import (
	"PismoAssessment/db"
	"PismoAssessment/errors"
	"PismoAssessment/models"
	"PismoAssessment/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

// CreateTransaction creates a new transaction.
// @Summary Create a new transaction
// @Description Initiates a transaction for a specified account, operation type, and amount.
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Transaction information"
// @Success 200 {object} models.Transaction "Successful transaction creation"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Router /v1/transactions [post]
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var trx models.Transaction
	err := json.NewDecoder(r.Body).Decode(&trx)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		logrus.Warn("Invalid request payload received for transaction")
		return
	}

	// Validate transaction fields
	if err := utils.Validate.Struct(trx); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := utils.FormatValidationErrors(validationErrors)
		errors.SendErrorResponse(w, http.StatusBadRequest, errorMessages)
		logrus.Warn("Validation failed for transaction creation:", errorMessages)
		return
	}

	err = db.DB.QueryRow("INSERT INTO pismo.transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, NOW()) RETURNING transaction_id", trx.AccountID, trx.OperationTypeID, trx.Amount).Scan(&trx.TransactionID)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create transaction")
		logrus.Error("Failed to create transaction:", err)
		return
	}

	// Now retrieve the full transaction data using the transaction_id
	err = db.DB.QueryRow(`
        SELECT transaction_id, account_id, operation_type_id, amount, event_date 
        FROM pismo.transactions 
        WHERE transaction_id = $1`, trx.TransactionID).
		Scan(&trx.TransactionID, &trx.AccountID, &trx.OperationTypeID, &trx.Amount, &trx.EventDate)

	if err != nil {
		errors.SendErrorResponse(w, http.StatusInternalServerError, "Failed to fetch transaction details")
		logrus.Error("Failed to fetch transaction details:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trx)
	logrus.Info("Transaction created successfully with ID:", trx.TransactionID)
}
