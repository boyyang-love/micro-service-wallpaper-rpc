package logic

import (
	"context"
	"email/helper"

	"email/internal/svc"
	"email/pb/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(in *email.SendEmailCodeReq) (*email.SendEmailCodeRes, error) {

	err := helper.SendEmail(helper.SendEmailParams{
		To:      in.Email,
		Subject: in.Subject,
		Code:    in.Code,
		Title:   in.Title,
		Time:    in.Time,
	})

	if err != nil {
		return nil, err
	}

	return &email.SendEmailCodeRes{
		Base: &email.Base{
			Code: 1,
			Msg:  "信息发送成功",
		},
	}, nil
}
