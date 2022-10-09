// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/rbacRole.proto

package rbacRole

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RbacRole service

func NewRbacRoleEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RbacRole service

type RbacRoleService interface {
	RoleGet(ctx context.Context, in *RoleGetRequest, opts ...client.CallOption) (*RoleGetResponse, error)
	RoleAdd(ctx context.Context, in *RoleAddRequest, opts ...client.CallOption) (*RoleAddResponse, error)
	RoleUpdate(ctx context.Context, in *RoleUpdateRequest, opts ...client.CallOption) (*RoleUpdateResponse, error)
	RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...client.CallOption) (*RoleDeleteResponse, error)
	RoleGetAuth(ctx context.Context, in *RoleGetAuthRequest, opts ...client.CallOption) (*RoleGetAuthResponse, error)
	RoleAuth(ctx context.Context, in *RoleAuthRequest, opts ...client.CallOption) (*RoleAuthResponse, error)
}

type rbacRoleService struct {
	c    client.Client
	name string
}

func NewRbacRoleService(name string, c client.Client) RbacRoleService {
	return &rbacRoleService{
		c:    c,
		name: name,
	}
}

func (c *rbacRoleService) RoleGet(ctx context.Context, in *RoleGetRequest, opts ...client.CallOption) (*RoleGetResponse, error) {
	req := c.c.NewRequest(c.name, "RbacRole.RoleGet", in)
	out := new(RoleGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacRoleService) RoleAdd(ctx context.Context, in *RoleAddRequest, opts ...client.CallOption) (*RoleAddResponse, error) {
	req := c.c.NewRequest(c.name, "RbacRole.RoleAdd", in)
	out := new(RoleAddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacRoleService) RoleUpdate(ctx context.Context, in *RoleUpdateRequest, opts ...client.CallOption) (*RoleUpdateResponse, error) {
	req := c.c.NewRequest(c.name, "RbacRole.RoleUpdate", in)
	out := new(RoleUpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacRoleService) RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...client.CallOption) (*RoleDeleteResponse, error) {
	req := c.c.NewRequest(c.name, "RbacRole.RoleDelete", in)
	out := new(RoleDeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacRoleService) RoleGetAuth(ctx context.Context, in *RoleGetAuthRequest, opts ...client.CallOption) (*RoleGetAuthResponse, error) {
	req := c.c.NewRequest(c.name, "RbacRole.RoleGetAuth", in)
	out := new(RoleGetAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacRoleService) RoleAuth(ctx context.Context, in *RoleAuthRequest, opts ...client.CallOption) (*RoleAuthResponse, error) {
	req := c.c.NewRequest(c.name, "RbacRole.RoleAuth", in)
	out := new(RoleAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RbacRole service

type RbacRoleHandler interface {
	RoleGet(context.Context, *RoleGetRequest, *RoleGetResponse) error
	RoleAdd(context.Context, *RoleAddRequest, *RoleAddResponse) error
	RoleUpdate(context.Context, *RoleUpdateRequest, *RoleUpdateResponse) error
	RoleDelete(context.Context, *RoleDeleteRequest, *RoleDeleteResponse) error
	RoleGetAuth(context.Context, *RoleGetAuthRequest, *RoleGetAuthResponse) error
	RoleAuth(context.Context, *RoleAuthRequest, *RoleAuthResponse) error
}

func RegisterRbacRoleHandler(s server.Server, hdlr RbacRoleHandler, opts ...server.HandlerOption) error {
	type rbacRole interface {
		RoleGet(ctx context.Context, in *RoleGetRequest, out *RoleGetResponse) error
		RoleAdd(ctx context.Context, in *RoleAddRequest, out *RoleAddResponse) error
		RoleUpdate(ctx context.Context, in *RoleUpdateRequest, out *RoleUpdateResponse) error
		RoleDelete(ctx context.Context, in *RoleDeleteRequest, out *RoleDeleteResponse) error
		RoleGetAuth(ctx context.Context, in *RoleGetAuthRequest, out *RoleGetAuthResponse) error
		RoleAuth(ctx context.Context, in *RoleAuthRequest, out *RoleAuthResponse) error
	}
	type RbacRole struct {
		rbacRole
	}
	h := &rbacRoleHandler{hdlr}
	return s.Handle(s.NewHandler(&RbacRole{h}, opts...))
}

type rbacRoleHandler struct {
	RbacRoleHandler
}

func (h *rbacRoleHandler) RoleGet(ctx context.Context, in *RoleGetRequest, out *RoleGetResponse) error {
	return h.RbacRoleHandler.RoleGet(ctx, in, out)
}

func (h *rbacRoleHandler) RoleAdd(ctx context.Context, in *RoleAddRequest, out *RoleAddResponse) error {
	return h.RbacRoleHandler.RoleAdd(ctx, in, out)
}

func (h *rbacRoleHandler) RoleUpdate(ctx context.Context, in *RoleUpdateRequest, out *RoleUpdateResponse) error {
	return h.RbacRoleHandler.RoleUpdate(ctx, in, out)
}

func (h *rbacRoleHandler) RoleDelete(ctx context.Context, in *RoleDeleteRequest, out *RoleDeleteResponse) error {
	return h.RbacRoleHandler.RoleDelete(ctx, in, out)
}

func (h *rbacRoleHandler) RoleGetAuth(ctx context.Context, in *RoleGetAuthRequest, out *RoleGetAuthResponse) error {
	return h.RbacRoleHandler.RoleGetAuth(ctx, in, out)
}

func (h *rbacRoleHandler) RoleAuth(ctx context.Context, in *RoleAuthRequest, out *RoleAuthResponse) error {
	return h.RbacRoleHandler.RoleAuth(ctx, in, out)
}
