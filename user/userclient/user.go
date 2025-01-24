// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: user.proto

package userclient

import (
	"context"

	"github.com/boyyang-love/micro-service-wallpaper-rpc/user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserReq = user.AddUserReq
	AddUserRes = user.AddUserRes
	Base       = user.Base

	User interface {
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserRes, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}
