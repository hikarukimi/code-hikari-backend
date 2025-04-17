package logic

import (
	"context"

	"code-hikari/user/api/internal/svc"
	"code-hikari/user/api/internal/types"
	"code-hikari/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type MobileLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMobileLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileLoginLogic {
	return &MobileLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MobileLoginLogic) MobileLogin(req *types.MobileLoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.MobileLogin(l.ctx, &service.MobileLoginRequest{
		Mobile:       req.Mobile,
		Verification: req.VerificationCode,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		UserId: res.UserId,
		Token:  res.Token,
	}, nil
}
