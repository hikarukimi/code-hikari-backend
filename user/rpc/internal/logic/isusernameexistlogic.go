package logic

import (
	"code-hikari/common-go"
	"context"

	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/server"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUsernameExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsUsernameExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUsernameExistLogic {
	return &IsUsernameExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsUsernameExistLogic) IsUsernameExist(in *server.IsUsernameExistRequest) (*server.IsUsernameExistResponse, error) {
	if l.svcCtx.Redis.BFExists(l.ctx, common.UsernameFilter, in.Username).Val() {
		return &server.IsUsernameExistResponse{
			Exist: true,
		}, nil
	}
	return &server.IsUsernameExistResponse{
		Exist: false,
	}, nil
}
