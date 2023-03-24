package transformer

import (
	"time"
)

type Format struct {
	Status    bool        `json:"status"`
	Service   string      `json:"service"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	CreatedAt int64       `json:"created_at"`
}

func Transformer(status bool, service string, message string, data interface{}) *Format {
	return &Format{
		Status:    status,
		Service:   service,
		Message:   message,
		Data:      data,
		CreatedAt: time.Now().Unix(),
	}
}
