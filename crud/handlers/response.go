package handlers

const (
	Error = "error"
	Message = "Message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(messsageType, message string, data interface{}) response {
	return response{
		messsageType,
		message,
		data,
	}
}


