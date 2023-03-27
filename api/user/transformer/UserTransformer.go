package transformer

import (
	"time"

	"github.com/irzam/my-app/api/user/utils"
	"github.com/irzam/my-app/lib/transformer"
)

type Format struct {
	Status     bool        `json:"status"`
	Service    string      `json:"service"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
	CreatedAt  int64       `json:"created_at"`
}

// Transformer with pagination
func Transformer(status bool, service string, message string, data interface{}, pagination interface{}) *Format {
	return &Format{
		Status:     status,
		Service:    service,
		Message:    message,
		Data:       data,
		Pagination: pagination,
		CreatedAt:  time.Now().Unix(),
	}
}

func UserGetAllTransformer(data map[string]interface{}) *Format {
	return Transformer(true, utils.Service, "Successfully", data["data"], data["pagination"])
}

func UserTransformer(data interface{}) *transformer.Format {
	return transformer.Transformer(true, utils.Service, "Successfully", data)
}
