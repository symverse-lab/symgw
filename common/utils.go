package common

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func GetBytes(value interface{}) ([]byte, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecodeBytes(value []byte, rvalue interface{}) error {
	err := json.Unmarshal(value, rvalue)
	if err != nil {
		return err
	}
	return nil
}

// My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors DynamicParameters `json:"error"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(DynamicParameters)

	switch err.(type) {
	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		for _, v := range errs {
			// can translate each error one at a time.
			if v.Param != "" {
				res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
			} else {
				res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
			}
		}
		break
	default:
		res.Errors["message"] = err.Error()
		break
	}

	return res
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(DynamicParameters)
	res.Errors[key] = err.Error()
	return res
}

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

func Hash(s interface{}) string {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)
	reqBodyBytes.Bytes()
	hash := md5.Sum([]byte(reqBodyBytes.Bytes()))
	return string(hash[:])
}
