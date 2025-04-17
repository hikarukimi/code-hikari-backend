package logic

import (
	"code-hikari/common-go"
	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/service"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsernameLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsernameLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsernameLoginLogic {
	return &UsernameLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsernameLoginLogic) UsernameLogin(in *service.UsernameLoginRequest) (*service.LoginResponse, error) {
	db := l.svcCtx.DB
	var loginUser common.User
	var notExistErr error
	notExistErr = db.First(&loginUser).Where("username=?", in.Username).Error

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
