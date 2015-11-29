package httpclient

type Network struct {
	message      string
	networkError error
}

func NewNetworkError(message string, networkError error) *Network {
	return &Network{message: message, networkError: networkError}
}

func (this *Network) Error() string {
	return this.message
}

func (this *Network) NetworkError() error {
	return this.networkError
}

func IsNetworkError(err interface{}) bool {
	_, ok := err.(*Network)
	return ok
}
