package main

import (
	"CSAwork/api"
	"CSAwork/boot"
	boot2 "CSAwork/boot/grpc"
)

func main() {
	boot.ViperSetup("./config/config.yaml")
	boot.LoggerSetup()
	boot.MysqlDBSetup()
	boot.RedisSetup()
	boot.MysqlSetUp()
	boot2.Grpcsetup()
	api.InitRouter()
}
