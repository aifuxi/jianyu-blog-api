package app

import (
	"github.com/aifuxi/jianyu-blog-api/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MakeErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}
