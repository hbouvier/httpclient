package httpclient

import "fmt"

type NotOk struct {
	message    string
	statusCode int
}

func NewNotOkError(message string, statusCode int) *NotOk {
	return &NotOk{message: message, statusCode: statusCode}
}

func (this *NotOk) Error() string {
	return fmt.Sprintf("[%3d] %s", this.statusCode, this.message)
}
func IsNotOkError(err interface{}) bool {
	_, ok := err.(*NotOk)
	return ok
}
