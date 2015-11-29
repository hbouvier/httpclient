package httpclient

type Unmarshalling struct {
	message            string
	unmarshallingError error
}

func NewUnmarshallingError(message string, unmarshallingError error) *Unmarshalling {
	return &Unmarshalling{message: message, unmarshallingError: unmarshallingError}
}

func (this *Unmarshalling) Error() string {
	return this.message
}

func (this *Unmarshalling) UnmarshallingError() error {
	return this.unmarshallingError
}

func IsUnmarshallingError(err interface{}) bool {
	_, ok := err.(*Unmarshalling)
	return ok
}
