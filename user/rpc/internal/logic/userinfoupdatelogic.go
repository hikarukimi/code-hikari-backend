package logic

import (
	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/server"
	"context"

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

func (l *UserInfoUpdateLogic) UserInfoUpdate(in *server.UserInfoUpdateRequest) (*server.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &server.UserInfoResponse{}, nil
}
