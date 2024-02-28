package exception

type OtherErrorException struct {
	Error string
}

func NewOtherError(error string) OtherErrorException  {
	return OtherErrorException{
		Error: error,
	}
}
