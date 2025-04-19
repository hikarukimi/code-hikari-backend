package logic

import (
	"code-hikari/user/rpc/server"
	"context"

	"code-hikari/user/api/internal/svc"
	"code-hikari/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUsernameExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsUsernameExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUsernameExistLogic {
	return &IsUsernameExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsUsernameExistLogic) IsUsernameExist(req *types.IsUsernameExistRequest) (resp *types.IsUsernameExistResponse, err error) {
	res, err := l.svcCtx.UserRpc.IsUsernameExist(l.ctx, &server.IsUsernameExistRequest{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	return &types.IsUsernameExistResponse{
		Exist: res.Exist,
	}, nil
}
