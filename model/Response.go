package model

// Response common response struct
type Response struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	AlreadySigned = -5003
)
