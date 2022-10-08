package models

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

const (
	consulAddrs = "127.0.0.1:8500"
)

var CaptchaClient client.Client

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
	CaptchaClient = srv.Client()

}
