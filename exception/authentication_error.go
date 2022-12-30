package exception

type AuthenticationError struct {
	Message string
}

func NewAuthenticationError(error string) AuthenticationError {
	return AuthenticationError{Message: error}
}

func (err AuthenticationError) Error() string {
	return err.Message
}
