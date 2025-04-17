package logic

import (
	"code-hikari/user/api/internal/svc"
	"code-hikari/user/api/internal/types"
	"code-hikari/user/rpc/service"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UsernameLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsernameLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsernameLoginLogic {
	return &UsernameLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsernameLoginLogic) UsernameLogin(req *types.UsernameLoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.UsernameLogin(l.ctx, &service.UsernameLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		UserId: res.UserId,
		Token:  res.Token,
	}, nil
}
