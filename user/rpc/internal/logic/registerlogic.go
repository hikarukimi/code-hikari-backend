package logic

import (
	"code-hikari/common-go"
	"code-hikari/user/rpc/internal/svc"
	"code-hikari/user/rpc/service"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {

	if l.svcCtx.Redis.Get(l.ctx, common.VerificationKeyPrefix+in.Mobile).Val() != in.VerificationCode {
		return nil, errors.New("验证码错误")
	}

	// 2. 检查用户是否已存在
	var existingUser common.User
	err := l.svcCtx.DB.Where("mobile = ? or username=?", in.Mobile, in.Username).First(&existingUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("数据库查询失败")
	}
	if existingUser.ID > 0 {
		return nil, errors.New("该用户已存在")
	}

	// 3. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密错误")
	}

	// 4. 创建新用户
	newUser := common.User{
		Username: in.Username,
		Mobile:   in.Mobile,
		Avatar:   *in.Avatar,
		Password: string(hashedPassword),
	}
	err = l.svcCtx.DB.Create(&newUser).Error
	if err != nil {
		return nil, errors.New("创建用户失败")
	}
	token, err := l.svcCtx.TokenUtil.CreateToken(int64(newUser.ID))
	if err != nil {
		return nil, errors.New("token生成错误")
	}

	// 5. 返回注册成功响应
	return &service.RegisterResponse{
		UserId: int64(newUser.ID),
		Token:  token,
	}, nil
}
