package logic

import (
	"context"
	"errors"
	"github.com/boyyang-love/micro-service-wallpaper-models/models"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/user/helper"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	"github.com/boyyang-love/micro-service-wallpaper-rpc/user/internal/svc"
	"github.com/boyyang-love/micro-service-wallpaper-rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRes, error) {
	password, err := helper.MakeHash(in.Password)
	if err != nil {
		return nil, err
	}
	if l.Is(in.Username) {
		userInfo := models.User{}
		if err := l.svcCtx.
			DB.
			Model(&models.User{}).
			Select("id", "username", "password").
			Where("username = ? and password = ?", in.Username, password).
			First(&userInfo).
			Error; errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}

		err, info := l.Info(userInfo.Id)
		if err != nil {
			return nil, err
		}

		token, err := helper.NewToken(&helper.JwtStruct{
			Id:               info.Id,
			Username:         info.Username,
			Role:             info.Role,
			RegisteredClaims: jwt.RegisteredClaims{},
		},
			l.svcCtx.Config.Token.AccessSecret,
			l.svcCtx.Config.Token.AccessExpire,
		)
		if err != nil {
			return nil, errors.New("token生成失败")
		}

		return &user.LoginRes{
			Base: &user.Base{
				Code: 1,
				Msg:  "登录成功",
			},
			Data: &user.LoginResData{
				Token: token,
				Info: &user.UserInfoResData{
					Id:       info.Id,
					Username: info.Username,
					Account:  info.Account,
					Motto:    info.Motto,
					Address:  info.Address,
					Tel:      info.Tel,
					Email:    info.Email,
					QQ:       info.QQ,
					Wechat:   info.Wechat,
					GitHub:   info.GitHub,
					Role:     info.Role,
					Avatar:   info.Avatar,
					Cover:    info.Cover,
				},
			},
		}, nil
	} else {
		return nil, errors.New("用户不存在")
	}
}

func (l *LoginLogic) Is(username string) (is bool) {
	userModel := models.User{}
	if err := l.svcCtx.
		DB.
		Model(&models.User{}).
		Select("username").
		Where("username = ?", username).
		First(&userModel).
		Error; errors.As(err, &gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (l *LoginLogic) Info(id string) (err error, info *models.User) {
	if err := l.svcCtx.
		DB.
		Model(&models.User{}).
		Select("Id", "Username", "Account", "Motto", "Address", "Tel", "Email", "QQ", "Wechat", "GitHub", "Role", "Avatar", "Cover").
		Where("id = ?", id).
		First(&info).
		Error; err != nil {
		return err, nil
	}

	return nil, info
}
