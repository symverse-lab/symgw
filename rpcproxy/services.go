package rpcproxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/symverse-lab/symgw/common"
	"github.com/symverse-lab/symgw/config"
	"github.com/symverse-lab/symgw/config/db"
	"net/http"
	"strconv"
)

var nonCachedRpcMethods = common.DynamicParameters{
	"sym_sendRawTransaction":    true,
	"sym_sendTransaction":       true,
	"sym_getTransactionReceipt": true,
	"sym_getTransactionCount":   true,
	"sym_getTransactionByHash":  true,
	"sct_getContract":           true,
	"sct_getContractAccount":    true,
	"sct_getAllowance":          true,
	"sct_getContractItem":       true,
}

// WorkNode 선택 및 Rpc Proxy API 요청
func (self *ProxyRouter) workNodeSelectAndRpcProxyCallService(c *gin.Context, request *common.JsonRpcRequest, number string) *common.DynamicParameters {
	workNodes := config.GetEnv().WorkNodes
	workNodeNumber, err := strconv.ParseInt(number, 10, 64)
	workNodeNumber = workNodeNumber - 1
	if len(workNodes) < int(workNodeNumber) || 0 > workNodeNumber {
		c.JSON(http.StatusBadRequest, common.NewError("message", fmt.Errorf("worknode가 존재하지 않습니다.")))
		return nil
	}

	workNodeHost := workNodes[workNodeNumber].HttpUrl
	if workNodeHost == "" || err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("message", fmt.Errorf("worknode가 존재하지 않습니다.")))
		return nil
	}

	_, isNonCacheMethod := nonCachedRpcMethods[request.Method]
	isCached := db.GetCache().IsCached(common.Hash(request))

	response, err := common.CachedRpcProxy(workNodeHost, c.Request, request, isCached && !isNonCacheMethod)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("message", err))
		return nil
	}
	return response
}

//WebSocket Rpc Proxy 처리
func (self *ProxyRouter) wsRpcProxyByNodeNumberService(c *gin.Context, request *common.JsonRpcRequest) interface{} {
	number := c.Param("node")
	response := self.workNodeSelectAndRpcProxyWsCallService(c, request, number)
	return response
}

// WorkNode 선택 및 WS Rpc Proxy API 요청
func (self *ProxyRouter) workNodeSelectAndRpcProxyWsCallService(c *gin.Context, request *common.JsonRpcRequest, number string) *common.DynamicParameters {
	workNodes := config.GetEnv().WorkNodes
	workNodeNumber, err := strconv.ParseInt(number, 10, 64)
	workNodeNumber = workNodeNumber - 1
	if len(workNodes) < int(workNodeNumber) || 0 > workNodeNumber {
		c.JSON(http.StatusBadRequest, common.NewError("message", fmt.Errorf("worknode가 존재하지 않습니다.")))
		return nil
	}
	workNodeHost := workNodes[workNodeNumber].WsUrl
	_, resbody, err := common.NodeRpcRequest(workNodeHost, request, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("message", err))
		return nil
	}
	return resbody
}
