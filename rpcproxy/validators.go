package rpcproxy

import (
	"github.com/gin-gonic/gin"
	"github.com/gocheat/symgw/common"
)

type JsonRpcRequestValidator struct {
	common.JsonRpcRequest
}

func (self *JsonRpcRequestValidator) Bind(c *gin.Context) (interface{}, error) {
	err := common.Bind(c, self)
	if err != nil {
		return nil, err
	}
	return self, nil
}

func NewRpcRequestValidator() *JsonRpcRequestValidator {
	jsonRpcModelValidator := &JsonRpcRequestValidator{}
	return jsonRpcModelValidator
}
