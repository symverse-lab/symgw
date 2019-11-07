package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/symverse-lab/symgw/bootnode"
	"github.com/symverse-lab/symgw/config"
	"github.com/symverse-lab/symgw/rpcproxy"
	"testing"

	_ "github.com/symverse-lab/symgw/docs"
)

// @Summary SYM-GW 애플리케이션의 상태를 체크합니다.
// @Description SYM-GW 애플리케이션의 상태를 체크합니다.
// @name Health Check
// @Accept  json
// @Produce  json
// @Router /health [get]
func TestRun(t *testing.T) {

	config.LoadEnvConfig("./env.yaml")
	db := databaseSelect()
	defer db.Close()

	r := gin.Default()
	r.Use(gin.Recovery())
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := r.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// rpcproxy api
	proxyRouter := rpcproxy.NewProxyRouter()
	bootNodeRpcProxyRouter := bootnode.NewBootnodeRpcProxyRouter()

	proxyRouter.RpcProxyRoute(v1.Group("/rpc"))
	bootNodeRpcProxyRouter.RpcProxyRoute(v1.Group("/bootnode"))

	httpListen(r)

}
