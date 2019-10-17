package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocheat/symgw/bootnode"
	"github.com/gocheat/symgw/config"
	"github.com/gocheat/symgw/config/db"
	"github.com/gocheat/symgw/rpcproxy"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gopkg.in/urfave/cli.v1"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/gocheat/symgw/docs"
)

const (
	VERSION  = "1.0.0"
	NAME     = "symgw"
	DESCRIBE = "symverse gateway server"
)

type Router interface {
	Stop() error
}

// @title SYM-GW API Docs
// @version 1.0.0
// @description Gsym API 문서입니다.
// @host testnet-gateway.symverse.com
func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = NAME
	app.Usage = DESCRIBE
	app.Flags = []cli.Flag{
		EnvFilePath,
		ModeFlag,
	}
	app.Action = func(ctx *cli.Context) error {
		run(ctx)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// @Summary SYM-GW 애플리케이션의 상태를 체크합니다.
// @Description SYM-GW 애플리케이션의 상태를 체크합니다.
// @name Health Check
// @Accept  json
// @Produce  json
// @Router /health [get]
func run(ctx *cli.Context) {
	//config file path
	envFile := ctx.GlobalString(EnvFilePath.Name)

	r := gin.Default()
	r.Use(gin.Recovery())

	//release Mode or dev Mode
	mode := ctx.GlobalString(ModeFlag.Name)
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)

		//file write log
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)

	} else {
		url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	config.LoadEnvConfig(envFile)
	db := databaseSelect()
	defer db.Close()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// REST API Route
	v1 := r.Group("/v1")
	proxyRouter := rpcproxy.NewProxyRouter()
	bootNodeRpcProxyRouter := bootnode.NewBootnodeRpcProxyRouter()

	proxyRouter.RpcProxyRoute(v1.Group("/rpc"))
	bootNodeRpcProxyRouter.RpcProxyRoute(v1.Group("/bootnode"))
	httpListen(r)
}

//databaseSelect
func databaseSelect() db.Cache {
	var (
		cache db.Cache
		err   error
	)

	if config.GetEnv().Cache.Use == false {
		return db.NewNonCache()
	}
	if config.GetEnv().Database.Driver == db.DB_LEVEL {
		cache, err = db.NewLevelDbCache("./_local")
	} else if config.GetEnv().Database.Driver == db.DB_REDIS {
		host := config.GetEnv().Database.Host
		if host == "" {
			host = DEFAULT_REDIS_HOST
		}
		port := config.GetEnv().Database.Port
		if port == "" {
			port = DEFAULT_REDIS_PORT
		}
		password := config.GetEnv().Database.Password
		cache, err = db.NewRedisCache(host+":"+port, password, 0)
	} else {
		log.Fatalf("데이터베이스 driver가 올바르지 않습니다.")
	}
	if err != nil {
		log.Fatalf("데이터베이스 접속 error 원인: %v", err)
	}
	return cache
}

//httpListen
func httpListen(r *gin.Engine) {
	addr := viper.Get("host.address")
	if addr == nil {
		addr = DEFAULT_HOST
	}
	port := viper.Get("host.port")
	if port == nil {
		port = DEFAULT_PORT
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", addr, port),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Http Listening => %v", srv.Addr)

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("GW Server exiting")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("Server exited")
}
