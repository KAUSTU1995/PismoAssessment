package controllers

import (
	"PismoAssessment/db"
	"PismoAssessment/errors"
	"PismoAssessment/models"
	"PismoAssessment/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// CreateAccount creates a new account.
// @Summary Create a new account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body models.Account true "Account information"
// @Success 200 {object} models.Account "Successful account creation"
// @Failure 400 {object} errors.ErrorResponse "Invalid input data"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Router /v1/accounts [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		logrus.Warn("Invalid request payload received")
		return
	}

	// Validate account fields
	if err := utils.Validate.Struct(acc); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := utils.FormatValidationErrors(validationErrors)
		errors.SendErrorResponse(w, http.StatusBadRequest, errorMessages)
		logrus.Warn("Validation failed for account creation:", errorMessages)
		return
	}

	err = db.DB.QueryRow("INSERT INTO pismo.accounts (document_number) VALUES ($1) RETURNING account_id", acc.DocumentNumber).Scan(&acc.AccountID)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create account")
		logrus.Error("Failed to create account:", err)
		return
	}

	// Respond with the created account ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
	logrus.Info("Account created successfully with ID:", acc.AccountID)
}

// GetAccount retrieves account information by ID.
// @Summary Get account information by ID
// @Description Fetches the account details for the specified account ID.
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Account "Successful response"
// @Failure 400 {object} errors.ErrorResponse "Invalid ID supplied"
// @Failure 404 {object} errors.ErrorResponse "Account not found"
// @Router /v1/accounts/{id} [get]
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Convert the ID from string to int
	accountID, err := strconv.Atoi(id)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusBadRequest, "Invalid account ID")
		logrus.Warn("Invalid account ID:", id)
		return
	}

	var acc models.Account
	err = db.DB.QueryRow("SELECT account_id, document_number FROM pismo.accounts WHERE account_id = $1", accountID).Scan(&acc.AccountID, &acc.DocumentNumber)
	if err != nil {
		errors.SendErrorResponse(w, http.StatusNotFound, "Account not found")
		logrus.Warn("Account not found with ID:", accountID)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
	logrus.Info("Account retrieved successfully with ID:", acc.AccountID)
}
