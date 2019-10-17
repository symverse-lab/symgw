package common

import (
	"github.com/gin-gonic/gin"
)

type DynamicParameters map[string]interface{}

type RpcWebSocketMessage struct {
	Parameters interface{}
	Context    *gin.Context
}

type JsonRpcRequest struct {
	JsonRpc string   `form:"jsonrpc" json:"jsonrpc" binding:"required",example:"2.0"`
	Id      uint     `form:"id" json:"id" binding:"required",example:"1"`
	Method  string   `form:"method" json:"method" binding:"required"`
	Params  []string `form:"params" json:"params" binding:"required"`
}

func (self *JsonRpcRequest) Bind(c *gin.Context) (interface{}, error) {
	err := Bind(c, self)
	if err != nil {
		return nil, err
	}
	return self, nil
}

func NewRpcRequest() *JsonRpcRequest {
	jsonRpcRequest := &JsonRpcRequest{}
	return jsonRpcRequest
}

type JsonRpcResponse struct {
	JsonRpc string      `json:"jsonrpc",example:"2.0"`
	Id      uint        `json:"id",example:"1"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error,omitempty"`
}
