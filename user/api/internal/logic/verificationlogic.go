package logic

import (
	"code-hikari/common-go"
	"code-hikari/user/api/internal/svc"
	"code-hikari/user/api/internal/types"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	mobile := req.Mobile
	redis := l.svcCtx.RedisClient
	result, err := redis.Get(l.ctx, mobile).Result()
	if err == nil {
		fmt.Println("redis中仍然存在验证码", err.Error())
	}
	if len(result) != 0 {
		return
	}
	s, err := redis.Set(l.ctx, common.VerificationKeyPrefix+mobile, "114514", 5*time.Minute).Result()
	if err != nil {
		fmt.Println("redis缓存失败", err.Error())
	}
	fmt.Println(s)
	return
}
