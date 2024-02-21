package exception

type PermissionError struct {
	Error string
}

func NewPermissionError(error string) PermissionError {
	return PermissionError{
		Error: error,
	}
}
