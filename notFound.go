package httpclient

type NotFound struct {
	message string
}

func NewNotFoundError(message string) *NotFound {
	return &NotFound{message: message}
}

func (this *NotFound) Error() string {
	return this.message
}
func IsNewNotFoundError(err interface{}) bool {
	_, ok := err.(*NotFound)
	return ok
}
