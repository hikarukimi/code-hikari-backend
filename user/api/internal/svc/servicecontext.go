package svc

import (
	"code-hikari/user/api/internal/config"
	"code-hikari/user/rpc/user"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
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
