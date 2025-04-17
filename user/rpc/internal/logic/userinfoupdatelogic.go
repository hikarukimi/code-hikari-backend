package logic

import (
	"context"

	"go-zhihu/application/user/internal/svc"
	"go-zhihu/application/user/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(in *service.UserInfoUpdateRequest) (*service.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &service.LoginResponse{}, nil
}
