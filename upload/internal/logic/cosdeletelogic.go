package logic

import (
	"context"
	"fmt"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/pb/upload"

	"github.com/zeromicro/go-zero/core/logx"
)

type CosDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCosDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CosDeleteLogic {
	return &CosDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CosDeleteLogic) CosDelete(in *upload.ImageDeleteReq) (*upload.Base, error) {
	for _, item := range in.Paths {
		imgPath := fmt.Sprintf("%s/%s", in.BucketName, item)
		err, _ := l.svcCtx.CosClient.Object.Delete(l.ctx, imgPath)
		if err != nil {
			panic(err)
		}
	}

	return &upload.Base{
		Code: 1,
		Msg:  "删除成功",
	}, nil
}
