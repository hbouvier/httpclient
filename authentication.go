package httpclient

type Authentication struct {
	message string
}

func NewAuthenticationError(message string) *Authentication {
	return &Authentication{message: message}
}

func (this *Authentication) Error() string {
	return this.message
}

func IsAuthenticationError(err interface{}) bool {
	_, ok := err.(*Authentication)
	return ok
}
