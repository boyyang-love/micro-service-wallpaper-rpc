package logic

import (
	"context"
	"errors"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/helper"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/models"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/upload/pb/upload"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadLogic) FileUpload(in *upload.FileUploadReq) (*upload.FileUploadRes, error) {

	is, info, err := l.IsHave(helper.MakeHashByBytes(in.File), in.FileType, in.UserId)
	if err != nil {
		return nil, err
	}

	if is {
		return &upload.FileUploadRes{
			Base: &upload.Base{
				Code: 1,
				Msg:  "文件上传成功",
			},
			Data: &upload.FileUploadResData{
				FileName:    in.FileName,
				OriFilePath: info.OriginFilePath,
				FilePath:    info.FilePath,
			},
		}, nil
	}

	compressedImage, err := helper.Image2Webp(in.File, uint(in.Quality))
	if err != nil {
		return nil, err
	}

	urls, err := helper.MinioFileUpload(&helper.MinioFileUploadParams{
		Ctx:         l.ctx,
		MinioClient: l.svcCtx.MinioClient,
		Buf:         compressedImage.Buf,
		OriBuf:      compressedImage.OriBuf,
		FileHash:    compressedImage.Hash,
		Folder:      in.Folder,
		OriFolder:   in.OriFolder,
		FileName:    in.FileName,
		FileType:    in.FileType,
		BucketName:  in.BucketName,
	})
	if err != nil {
		return nil, err
	}

	if err = l.InsertToMySql(&models.Upload{
		Hash:           compressedImage.Hash,
		FileName:       in.FileName,
		OriginFileSize: compressedImage.OriSize,
		FileSize:       compressedImage.Size,
		OriginType:     in.FileType,
		FileType:       "webp",
		OriginFilePath: urls.OriFilePath,
		FilePath:       urls.FilePath,
		UserId:         in.UserId,
		Type:           in.Type,
		Status:         false,
		W:              int(in.W),
		H:              int(in.H),
	}); err != nil {
		return nil, err
	}

	return &upload.FileUploadRes{
		Base: &upload.Base{
			Code: 1,
			Msg:  "文件上传成功",
		},
		Data: &upload.FileUploadResData{
			FileName:    in.FileName,
			OriFilePath: urls.OriFilePath,
			FilePath:    urls.FilePath,
		},
	}, nil
}

func (l *FileUploadLogic) InsertToMySql(uploadParams *models.Upload) error {
	if err := l.svcCtx.
		DB.
		Model(&models.Upload{}).
		Create(uploadParams).
		Error; err != nil {
		return err
	}

	return nil
}

func (l *FileUploadLogic) IsHave(hash string, fileType string, userId string) (is bool, info *models.Upload, err error) {
	if err := l.svcCtx.
		DB.
		Model(&models.Upload{}).
		Select("hash", "type", "user_id", "file_name", "file_path", "origin_file_path").
		Where("hash = ? and type = ? and user_id = ?", hash, fileType, userId).
		First(&info).
		Error; errors.As(err, &gorm.ErrRecordNotFound) {
		return false, info, nil
	} else {
		return true, info, nil
	}
}
