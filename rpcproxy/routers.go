package rpcproxy

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocheat/symgw/common"
	"github.com/gocheat/symgw/config"
	"github.com/gocheat/symgw/config/db"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type ProxyRouter struct {
	cacheSaveEvent chan common.CacheSaveEvent
	done           chan bool
}

func NewProxyRouter() *ProxyRouter {
	router := &ProxyRouter{
		cacheSaveEvent: make(chan common.CacheSaveEvent),
		done:           make(chan bool),
	}
	go router.eventLoop()
	return router
}

// 1. 파라미터 검증
// 2. 노드 heath active 체크
// 3. RPC Method 체크
// 3-1. read 함수일경우 캐시레이어에 저장 후 node proxy
// 3-2. write 홤수일경우 direct node proxy
func (self *ProxyRouter) RpcProxyRoute(router *gin.RouterGroup) {

	router.POST("/node/:node", self.httpRpcProxyByNodeNumber)
	//router.GET("/node/:node/ws", self.wsRpcHandler)
	router.GET("/nodes", self.getStaticWorkNodes)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (self *ProxyRouter) Stop() error {
	<-self.done
	return nil
}

func (self *ProxyRouter) eventLoop() {

	ticker := time.NewTicker(5 * time.Second)
	//health check
	for {
		select {
		case <-self.done:
			return
		case <-ticker.C:
			db.GetCache().DeleteExpired()
		}
	}
}

// @Summary WorkNode로 WebSocket JSON-RPC 요청을 Proxy 합니다.
// @Description WorkNode로 HTTP JSON-RPC 요청을 Proxy 합니다.
// @name httpRpcProxyByNodeNumber
// @Accept  json
// @Produce  json
// @Param name path string true "Node Number"
// @Param message body common.JsonRpcRequest true "Json-Rpc Request"
// @Router /v1/rpc/node/{number} [post]
// @Success 200 {object} common.JsonRpcResponse
func (self *ProxyRouter) httpRpcProxyByNodeNumber(c *gin.Context) {
	request := common.NewRpcRequest()
	_, err := request.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}
	number := c.Param("node")
	response := self.workNodeSelectAndRpcProxyCallService(c, request, number)
	c.JSON(http.StatusOK, response)
}

// @Summary WorkNode로 WebSocket JSON-RPC 요청을 Proxy 합니다.
// @Description WorkNode로 WebSocket JSON-RPC 요청을 Proxy 합니다.
// @name wsRpcProxyByNodeNumber
// @Accept  json
// @Produce  json
// @Param name path string true "Node Number"
// @Param message body common.JsonRpcRequest true "Json-Rpc Request"
// @Router /v1/rpc/node/{number}/ws [post]
// @Success 200 {object} common.JsonRpcResponse
//WebSocket Handler
func (self *ProxyRouter) wsRpcHandler(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("message", errors.New(fmt.Sprintf("Failed to set websocket upgrade: %+v", err))))
		return
	}
	defer conn.Close()

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		params := &common.JsonRpcRequest{}
		err = common.DecodeBytes(msg, params)
		if err != nil {
			errMessage, _ := common.GetBytes(common.NewError("message", err))
			conn.WriteMessage(t, errMessage)
		} else {
			responseBytes := self.wsRpcProxyByNodeNumberService(c, params)
			response, _ := common.GetBytes(responseBytes)
			conn.WriteMessage(t, response)
		}
		//
	}
}

// @Summary Get Static WorkNodes
// @Description Proxy를 제공할 URL 정보가 있는 WorkNode 리스트를 요청합니다.
// @name wsRpcProxyByNodeNumber
// @Accept  json
// @Produce  json
// @Router /v1/rpc/nodes [get]
// @Success 200 {array} config.WorkNode
func (self *ProxyRouter) getStaticWorkNodes(c *gin.Context) {
	workNodes := config.GetEnv().WorkNodes
	for i, workNode := range workNodes {
		workNode.Number = i + 1
	}
	c.JSON(http.StatusOK, workNodes)
}

///////////////////////////////
