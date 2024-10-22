
CREATE SCHEMA IF NOT EXISTS pismo;

CREATE TABLE IF NOT EXISTS pismo.accounts (
                                              account_id SERIAL PRIMARY KEY,
                                              document_number VARCHAR(50) NOT NULL
    );

CREATE TABLE IF NOT EXISTS pismo.transactions (
                                                  transaction_id SERIAL PRIMARY KEY,
                                                  account_id INT REFERENCES pismo.accounts(account_id),
    operation_type_id INT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    event_date TIMESTAMP NOT NULL DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS pismo.operation_types (
                                                     operation_type_id INT PRIMARY KEY,
                                                     description VARCHAR(50) NOT NULL
    );
