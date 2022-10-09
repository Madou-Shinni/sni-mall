package models

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

var RbacClient client.Client

// init 注册进consul
func init() {
	//配置consul
	consulReg := consul.NewRegistry(
		registry.Addrs(consulAddrs),
	)
	// Create service
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	srv.Init()
	RbacClient = srv.Client()

}
