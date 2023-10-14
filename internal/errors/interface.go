package errors

// Encapsulates errors to be returned via API
type ExternalError interface {
	Error() string
	HTTPCode() int
}
