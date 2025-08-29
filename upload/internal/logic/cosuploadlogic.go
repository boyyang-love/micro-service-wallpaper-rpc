package logic

import (
	"context"
	"fmt"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/helper"
	"strings"

	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/pb/upload"

	"github.com/zeromicro/go-zero/core/logx"
)

type CosUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCosUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CosUploadLogic {
	return &CosUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CosUploadLogic) CosUpload(in *upload.ImageUploadReq) (*upload.ImageUploadRes, error) {
	comp, err := helper.Image2Webp(&in.File, uint(in.Quality))
	if err != nil {
		return nil, err
	}

	imgOriPath := fmt.Sprintf("%s/%s", in.BucketName, in.OriPath)
	imgComPath := fmt.Sprintf("%s/%s", in.BucketName, in.Path)

	_, err = l.
		svcCtx.
		CosClient.
		Object.
		Put(l.ctx, imgComPath, comp.Buf, nil)
	if err != nil {
		return nil, err
	}

	_, err = l.
		svcCtx.
		CosClient.
		Object.
		Put(l.ctx, imgOriPath, comp.OriBuf, nil)
	if err != nil {
		return nil, err
	}

	oriInfo, err := l.svcCtx.CosClient.Object.Head(l.ctx, imgOriPath, nil)
	if err != nil {
		return nil, err
	}

	comInfo, err := l.svcCtx.CosClient.Object.Head(l.ctx, imgComPath, nil)
	if err != nil {
		return nil, err
	}

	return &upload.ImageUploadRes{
		Base: &upload.Base{
			Code: 1,
			Msg:  "图片上传成功",
		},
		Data: &upload.ImageUploadResData{
			Path:    in.Path,
			OriPath: in.OriPath,
			ETag:    strings.ReplaceAll(comInfo.Header.Get("ETag"), "\"", ""),
			OriETag: strings.ReplaceAll(oriInfo.Header.Get("ETag"), "\"", ""),
			Size:    uint64(comp.Size),
			OriSize: uint64(comp.OriSize),
		},
	}, nil
}
