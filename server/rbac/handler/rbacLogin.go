package handler

import (
	"context"
	"rbac/models"
	mysql "rbac/models/mysql"
	"rbac/models/utils"
	pb "rbac/proto/rbacLogin"
)

type RbacLogin struct{}

func (e *RbacLogin) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	var manager models.Manager
	affected := mysql.DB.Where("username = ? And password = ?", req.Username, req.Password).Find(&manager).RowsAffected
	if affected > 0 {
		token, err := utils.GenToken(manager.Id, manager.Username)
		if err != nil {
			return err
		}
		rsp.Token = token
	}
	return nil
}
