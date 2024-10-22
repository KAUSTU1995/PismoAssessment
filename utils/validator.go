package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitializeValidator() {
	Validate = validator.New()

	// Register custom validation function
	Validate.RegisterValidation("operationtype", operationTypeValidator)
}

func FormatValidationErrors(validationErrors validator.ValidationErrors) string {
	var errMessages string
	for _, err := range validationErrors {
		errMessages += "Field '" + err.Field() + "' with value '" + fmt.Sprintf("%v", err.Value()) + "' failed validation, reason: '" + err.Tag() + "'. "
	}
	return errMessages
}

// Custom validation function to check if the OperationTypeID is valid
func operationTypeValidator(fl validator.FieldLevel) bool {
	operationTypeID := fl.Field().Int()
	validOperationTypes := []int{1, 2, 3, 4} // Allowed Operation Type IDs
	for _, id := range validOperationTypes {
		if operationTypeID == int64(id) {
			return true
		}
	}
	return false
}
