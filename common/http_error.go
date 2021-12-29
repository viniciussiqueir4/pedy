package common

type HttpError struct {
	StatusCode int
	Errors []error
}
var BaseNotFoundText = "No record found for provided parameter."
var BaseNoRecordAffected = "None records were changed."
