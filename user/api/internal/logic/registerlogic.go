package logic

import (
	"code-hikari/common-go"
	"code-hikari/user/api/internal/svc"
	"code-hikari/user/api/internal/types"
	"code-hikari/user/rpc/server"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	result, err := l.svcCtx.UserRpc.Register(l.ctx, &server.RegisterRequest{
		Username:         req.Username,
		Mobile:           req.Mobile,
		Password:         req.Password,
		VerificationCode: req.VerificationCode,
	})
	if err != nil {
		return nil, errors.New("用户创建失败" + err.Error())
	}
	// 将用户名添加到布隆过滤器中
	l.svcCtx.RedisClient.BFAdd(l.ctx, common.UsernameFilter, req.Username)
	return &types.RegisterResponse{
		UserId: result.UserId,
		Token:  result.Token,
	}, nil
}
