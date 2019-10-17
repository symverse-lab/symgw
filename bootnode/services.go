package bootnode

import (
	"github.com/gin-gonic/gin"
	"github.com/gocheat/symgw/common"
	"github.com/gocheat/symgw/config"
	"github.com/gocheat/symgw/config/db"
	"net/http"
)

// BootNode RPC API Proxy
func (self *BootNodeProxyRouter) getNodesByBootNodeRpcProxyCallService(c *gin.Context, rpcMethod string) *common.DynamicParameters {

	request := common.NewRpcRequest()
	request.Id = 1
	request.Method = rpcMethod
	request.JsonRpc = "2.0"

	bootNodeUrl := config.GetEnv().BootNodes[0].HttpUrl

	var response *common.DynamicParameters
	isCached := db.GetCache().IsCached(rpcMethod)
	response, err := common.CachedRpcProxy(bootNodeUrl, c.Request, request, isCached)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("message", err))
		return nil
	}

	return response
}
