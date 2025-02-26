package logic

import (
	"context"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/pb/upload"
	"github.com/minio/minio-go/v7"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageDeleteLogic {
	return &ImageDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ImageDeleteLogic) ImageDelete(in *upload.ImageDeleteReq) (*upload.Base, error) {

	for _, item := range in.Paths {
		err := l.svcCtx.MinioClient.RemoveObject(l.ctx, in.BucketName, item, minio.RemoveObjectOptions{})
		if err != nil {
			panic(err)
		}
	}

	return &upload.Base{
		Code: 1,
		Msg:  "删除成功",
	}, nil
}
