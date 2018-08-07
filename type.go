package beescaffold

import "github.com/pkg/errors"

var DefErrors map[error]int
var DefaultCode int =-1

type R struct {
	Error int `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type E struct {
	Error error
	Code int
}

var E_NotLogin=Make_E_with_s_code("尚未登录",-2)

func Api_R_with_obj(_o interface{})*R{
	return &R{Data:_o}
}

func Api_Error_with_error_code(_err error,_n int)*R{
	return &R{Error:_n,Message:_err.Error()}
}

func Api_Error_with_e(_e *E)*R{
	return Api_Error_with_error_code(_e.Error,_e.Code)
}

func Make_E_defaultcode_with_err(_err error)*E{
	return Make_E_with_error_code(_err,DefaultCode)
}

func Make_E_defaultcode_with_s(_s string)*E{
	return Make_E_with_s_code(_s,DefaultCode)
}

func Make_E_with_s_code(_s string,_code int)*E{
	return Make_E_with_error_code(errors.New(_s),_code)
}



func Make_E_with_error_code(_err error,_code int)*E{
	return &E{Error:_err,Code:_code}
}