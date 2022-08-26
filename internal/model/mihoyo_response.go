package model

// MihoyoResponse common response struct
type MihoyoResponse struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	AlreadySigned = -5003
)
