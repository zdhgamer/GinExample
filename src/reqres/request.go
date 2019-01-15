package reqres

import (
	"github.com/astaxie/beego/validation"
	"../logging"
)

func MarkError(errors []*validation.Error)  {
	for _,err:= range errors {
		logging.Info(err.Key,err.Value)
	}
	return
}
