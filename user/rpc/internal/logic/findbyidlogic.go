package logic

import (
	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/server"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *server.FindByIdRequest) (*server.FindByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &server.FindByIdResponse{}, nil
}
