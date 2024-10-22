package main

import (
	"PismoAssessment/controllers"
	"PismoAssessment/db"
	"PismoAssessment/models"
	"PismoAssessment/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

var (
	mockDB *sql.DB
	mock   sqlmock.Sqlmock
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.SetLevel(logrus.DebugLevel)
}

func TestMain(m *testing.M) {
	var err error
	mockDB, mock, err = sqlmock.New()
	if err != nil {
		logrus.Fatal("Failed to open mock database:", err)
	}
	db.DB = mockDB
	logrus.Info("Starting test suite")

	utils.InitializeValidator()

	m.Run()
}

func TestCreateAccount(t *testing.T) {
	logrus.Info("Running TestCreateAccount")

	reqBody := `{"document_number":"12345678900"}`
	req, err := http.NewRequest("POST", "/v1/accounts", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set up  the mock database
	mock.ExpectQuery(`^INSERT INTO pismo\.accounts \(document_number\) VALUES \(\$1\) RETURNING account_id$`).
		WithArgs("12345678900").
		WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow(99))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateAccount)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var acc models.Account
	err = json.NewDecoder(rr.Body).Decode(&acc)
	if err != nil || acc.AccountID == 0 {
		t.Errorf("Expected valid account creation, got %v", rr.Body.String())
	}
}

// Test for getting an account by ID
func TestGetAccount(t *testing.T) {
	logrus.Info("Running TestGetAccount")

	accountID := 1
	req, err := http.NewRequest("GET", "/v1/accounts/"+strconv.Itoa(accountID), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up expected behavior for the mock database
	mock.ExpectQuery(`^SELECT account_id, document_number FROM pismo\.accounts WHERE account_id = \$1$`).
		WithArgs(accountID).
		WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow(accountID, "12345678900"))

	rr := httptest.NewRecorder()
	mux := mux.NewRouter()
	mux.HandleFunc("/v1/accounts/{id}", controllers.GetAccount)

	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var acc models.Account
	err = json.NewDecoder(rr.Body).Decode(&acc)
	if err != nil || acc.AccountID == 0 {
		t.Errorf("Expected valid account retrieval, got %v", rr.Body.String())
	}
}

func TestCreateTransaction(t *testing.T) {
	logrus.Info("Running TestCreateTransaction")

	// Request body to test
	reqBody := `{"account_id":1,"operation_type_id":2,"amount":100.50}`
	req, err := http.NewRequest("POST", "/v1/transactions", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set up the mock database insert query expectation
	mock.ExpectQuery(`^INSERT INTO pismo.transactions \(account_id, operation_type_id, amount, event_date\) VALUES \(\$1, \$2, \$3, NOW\(\)\) RETURNING transaction_id$`).
		WithArgs(1, 2, 100.50).
		WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow(1))

	// Set up the mock database select query expectation for fetching transaction details
	mock.ExpectQuery(`^SELECT transaction_id, account_id, operation_type_id, amount, event_date FROM pismo.transactions WHERE transaction_id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"transaction_id", "account_id", "operation_type_id", "amount", "event_date"}).
			AddRow(1, 1, 2, 100.50, time.Now()))

	// Prepare the response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateTransaction)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body into a transaction struct
	var trx models.Transaction
	err = json.NewDecoder(rr.Body).Decode(&trx)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Validate the returned transaction
	if trx.TransactionID == 0 {
		t.Errorf("Expected valid transaction creation, got %v", rr.Body.String())
	}

	// Ensure all expectations are met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}
