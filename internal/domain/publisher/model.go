package publisher

type PublishRequest struct {
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
}
