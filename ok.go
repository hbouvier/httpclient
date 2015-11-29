package httpclient

var OK error = nil

func NewOkError() error {
	return OK
}

func isOK(err interface{}) bool {
	if err == nil {
		return true
	} else {
		return false
	}
}
