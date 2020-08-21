package backend_errors

type ErrorType string

const (
	Email    ErrorType = "email"
	Required ErrorType = "required"
)

var ErrorTypes = map[ErrorType]string{
	Email:    "The email given is not correct",
	Required: "The field is required",
}
