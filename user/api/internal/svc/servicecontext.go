package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zhihu/application/applet/internal/config"
	"go-zhihu/application/user/user"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	UserRpc     user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	return &ServiceContext{
		Config:      c,
		RedisClient: rdb,
		UserRpc:     user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
