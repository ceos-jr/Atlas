package utils

import "fmt"

type CustomError struct {
	Message string
}

func NewError(errorLabel string, referenceError error) *CustomError {
  return &CustomError{
		Message: fmt.Sprintf("%s: %v", errorLabel, referenceError),
	}
}

func (c *CustomError) AddLabel(adicionalLabel string) {
	c.Message = fmt.Sprintf("%s %s", adicionalLabel, c.Message)
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Error: %s ", c.Message)
}
