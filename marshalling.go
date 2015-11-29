package httpclient

type Marshalling struct {
	message          string
	marshallingError error
}

func NewMarshallingError(message string, marshallingError error) *Marshalling {
	return &Marshalling{message: message, marshallingError: marshallingError}
}

func (this *Marshalling) Error() string {
	return this.message
}

func (this *Marshalling) MarshallingError() error {
	return this.marshallingError
}

func IsMarshallingError(err interface{}) bool {
	_, ok := err.(*Marshalling)
	return ok
}
