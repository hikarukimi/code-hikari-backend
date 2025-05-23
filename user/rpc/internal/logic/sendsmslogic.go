package logic

import (
	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/server"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *server.SendSmsRequest) (*server.SendSmsResponse, error) {
	// todo: add your logic here and delete this line

	return &server.SendSmsResponse{}, nil
}
