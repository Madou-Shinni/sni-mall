package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"rbac/handler"
	. "rbac/models/ini"
	pbLogin "rbac/proto/rbacLogin"
	pbManager "rbac/proto/rbacManager"
	pbRole "rbac/proto/rbacRole"
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
	// 注册登录
	if err := pbLogin.RegisterRbacLoginHandler(srv.Server(), new(handler.RbacLogin)); err != nil {
		logger.Fatal(err)
	}
	// 注册角色管理
	if err := pbRole.RegisterRbacRoleHandler(srv.Server(), new(handler.RbacRole)); err != nil {
		logger.Fatal(err)
	}
	// 注册管理员管理
	if err := pbManager.RegisterRbacManagerHandler(srv.Server(), new(handler.RbacManager)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
