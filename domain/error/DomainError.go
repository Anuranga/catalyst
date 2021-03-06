package error

import "fmt"

// DomainError is the type of errors thrown by business logic.
type DomainError struct {
	msg     string
	code    int
	details string
}

// New creates a new DomainError instance.
func (e *DomainError) New(message string, code int, details string) error {

	err := &DomainError{}

	err.msg = message
	err.code = code
	err.details = details

	return err
}

// Error returns the DomainError message.
func (e *DomainError) Error() string {
	return fmt.Sprintf("%s|%d|DomainError|%s", e.msg, e.code, e.details)
}
