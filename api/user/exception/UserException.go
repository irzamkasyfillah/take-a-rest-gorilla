package exception

import (
	"github.com/irzam/my-app/api/user/utils"
	"github.com/irzam/my-app/lib/transformer"
)

type HandleError struct {
	Message    string
	Data       map[string]interface{}
	StatusCode int
}

func UserException(message string, data interface{}) *transformer.Format {
	return transformer.Transformer(false, utils.Service, message, data)
}
