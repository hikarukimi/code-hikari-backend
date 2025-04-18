package logic

import (
	"code-hikari/common-go/model"
	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/server"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsernameLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsernameLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsernameLoginLogic {
	return &UsernameLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsernameLoginLogic) UsernameLogin(in *server.UsernameLoginRequest) (*server.LoginResponse, error) {
	db := l.svcCtx.DB
	var loginUser model.User
	var notExistErr error
	notExistErr = db.Where("username=?", in.Username).First(&loginUser).Error
	fmt.Println(loginUser)
	if notExistErr != nil {
		return nil, errors.New("不存在该用户")
	}
	token, err := l.svcCtx.TokenUtil.CreateToken(int64(loginUser.ID))
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	return &server.LoginResponse{
		Token:  token,
		UserId: int64(loginUser.ID),
	}, nil
}
