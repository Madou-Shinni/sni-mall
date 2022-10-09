package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"rbac/handler"
	. "rbac/models/ini"
	pb "rbac/proto/rbac"
)

var (
	service = "rbac"
	version = "latest"
)

func main() {
	// 读取ini配置
	addr := Config.Section("consul").Key("addr").String()
	// 注册consul
	consulRegistry := consul.NewRegistry(registry.Addrs(addr))
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulRegistry),
	)

	// Register handler
	if err := pb.RegisterRbacHandler(srv.Server(), new(handler.Rbac)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
