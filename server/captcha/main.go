package main

import (
	"captcha/handler"
	pb "captcha/proto/captcha"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service     = "captcha"
	version     = "latest"
	consulAddrs = "127.0.0.1:8500"
)

func main() {
	// 配置consul
	consulRegistry := consul.NewRegistry(registry.Addrs(consulAddrs))
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulRegistry),
	)

	// Register handler
	if err := pb.RegisterCaptchaHandler(srv.Server(), new(handler.Captcha)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
