// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: email.proto

package server

import (
	"context"

	"github.com/boyyang-love/micro-service-wallpaper-rpc/email/internal/logic"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/email/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/email/pb/email"
)

type EmailServer struct {
	svcCtx *svc.ServiceContext
	email.UnimplementedEmailServer
}

func NewEmailServer(svcCtx *svc.ServiceContext) *EmailServer {
	return &EmailServer{
		svcCtx: svcCtx,
	}
}

func (s *EmailServer) SendEmail(ctx context.Context, in *email.SendEmailReq) (*email.SendEmailRes, error) {
	l := logic.NewSendEmailLogic(ctx, s.svcCtx)
	return l.SendEmail(in)
}

func (s *EmailServer) SendEmailCode(ctx context.Context, in *email.SendEmailCodeReq) (*email.SendEmailCodeRes, error) {
	l := logic.NewSendEmailCodeLogic(ctx, s.svcCtx)
	return l.SendEmailCode(in)
}
