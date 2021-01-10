package movieerror

type movieError struct {
}

func (e *movieError) Error() string {
	return ""
}

var MovieNotFound = &movieError{
}
