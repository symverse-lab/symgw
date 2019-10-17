package bootnode

import (
	"github.com/gin-gonic/gin"
	"github.com/gocheat/symgw/common"
	"net/http"
	"time"
)

type BootNodeProxyRouter struct {
	cacheSaveEvent chan common.CacheSaveEvent
	done           chan bool
}

func NewBootnodeRpcProxyRouter() *BootNodeProxyRouter {
	router := &BootNodeProxyRouter{
		cacheSaveEvent: make(chan common.CacheSaveEvent),
		done:           make(chan bool),
	}
	go router.eventLoop()
	return router
}

func (self *BootNodeProxyRouter) RpcProxyRoute(router *gin.RouterGroup) {
	router.GET("/nodes", self.GetNodesByBootNodeRpcProxy)
	router.GET("/closeNodes", self.GetCloseNodesByBootNodeRpcProxy)
}

func (self *BootNodeProxyRouter) Stop() error {
	<-self.done
	return nil
}

func (self *BootNodeProxyRouter) eventLoop() {

	ticker := time.NewTicker(5 * time.Second)
	//health check
	for {
		select {
		case <-self.done:
			return
		case <-ticker.C:
		}
	}
}

// @Summary BootNode API 요청하여 Node 정보를 제공합니다.
// @Description BootNode API 요청하여 Node 정보를 제공합니다.
// @name Call BootNode Nodes Api
// @Accept  json
// @Produce  json
// @Router /v1/bootnode/nodes [get]
// @Success 200 {object} common.JsonRpcResponse
func (self *BootNodeProxyRouter) GetNodesByBootNodeRpcProxy(c *gin.Context) {
	response := self.getNodesByBootNodeRpcProxyCallService(c, "bootnode_getNodes")
	c.JSON(http.StatusOK, response)
}

// @Summary BootNode API 요청하여 가장 근접한 Node 정보를 제공합니다.
// @Description BootNode API 요청하여 Node 정보를 제공합니다.
// @name Call BootNode closeNodes Api
// @Accept  json
// @Produce  json
// @Router /v1/bootnode/closeNodes [get]
// @Success 200 {object} common.JsonRpcResponse
func (self *BootNodeProxyRouter) GetCloseNodesByBootNodeRpcProxy(c *gin.Context) {
	response := self.getNodesByBootNodeRpcProxyCallService(c, "bootnode_getCloseNodes")
	c.JSON(http.StatusOK, response)
}
