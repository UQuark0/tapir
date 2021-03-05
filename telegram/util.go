package telegram

type ResponseError struct {
	error
	responseMessage string
}

func NewResponseError(message string, err error) *ResponseError {
	re := ResponseError{}
	if err != nil {
		re.error = err
	}
	re.responseMessage = message
	return &re
}

func (re *ResponseError) ResponseMessage() string {
	return re.responseMessage
}