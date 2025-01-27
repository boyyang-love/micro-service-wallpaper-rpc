// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: user.proto

package server

import (
	"context"

	"github.com/boyyang-love/micro-service-wallpaper-rpc/user/internal/logic"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/user/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/user/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) AddUser(ctx context.Context, in *user.AddUserReq) (*user.AddUserRes, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}
