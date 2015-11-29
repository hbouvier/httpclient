package httpclient

import "fmt"

type ServerFailure struct {
	message    string
	statusCode int
}

func NewServerFailureError(message string, statusCode int) *ServerFailure {
	return &ServerFailure{message: message, statusCode: statusCode}
}

func (this *ServerFailure) Error() string {
	return fmt.Sprintf("[%3d] %s", this.statusCode, this.message)
}

func IsServerFailureError(err interface{}) bool {
	_, ok := err.(*ServerFailure)
	return ok
}
