package main

import (
	cli "gopkg.in/urfave/cli.v1"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = "8080"

	DEFAULT_REDIS_HOST = "localhost"
	DEFAULT_REDIS_PORT = "6379"
)

var (
	EnvFilePath = cli.StringFlag{
		Name:  "env",
		Usage: "Load configuration from `FILE`",
		Value: "env.yaml",
	}

	ModeFlag = cli.StringFlag{
		Name:  "mode",
		Usage: "default `debug` mode. Switch to `release` mode in production.",
		Value: "debug",
	}
)
