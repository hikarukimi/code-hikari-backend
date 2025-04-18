package logic

import (
	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/server"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *server.FindByMobileRequest) (*server.FindByMobileResponse, error) {
	// todo: add your logic here and delete this line

	return &server.FindByMobileResponse{}, nil
}
