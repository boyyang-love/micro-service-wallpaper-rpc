package logic

import (
	"context"

	"email/internal/svc"
	"email/pb/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *email.SendEmailReq) (*email.SendEmailRes, error) {
	// todo: add your logic here and delete this line

	return &email.SendEmailRes{}, nil
}
