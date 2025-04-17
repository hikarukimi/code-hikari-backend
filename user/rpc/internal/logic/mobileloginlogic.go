package logic

import (
	"context"
	"errors"
	"go-zhihu/common"

	"go-zhihu/application/user/internal/svc"
	"go-zhihu/application/user/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type MobileLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMobileLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileLoginLogic {
	return &MobileLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MobileLoginLogic) MobileLogin(in *service.MobileLoginRequest) (*service.LoginResponse, error) {
	db := l.svcCtx.DB
	var loginUser common.User
	var notExistErr error

	notExistErr = db.First(&loginUser).Where("mobile=?", in.Mobile).Error
	if notExistErr != nil {
		return nil, errors.New("不存在该用户")
	}
	token, err := l.svcCtx.TokenUtil.CreateToken(int64(loginUser.ID))
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	return &service.LoginResponse{
		Token:  token,
		UserId: int64(loginUser.ID),
	}, nil
}
