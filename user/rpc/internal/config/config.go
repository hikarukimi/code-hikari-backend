package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

// Config 定义主配置结构体
type Config struct {
	zrpc.RpcServerConf // 匿名嵌套 RpcServerConf
	PostgresqlConfig   PostgresqlConfig
	TokenAuth          AuthConfig // 认证配置
	RedisConfig        redis.RedisConf
	Consul             consul.Conf
}

// AuthConfig 定义认证相关的配置
type AuthConfig struct {
	AccessSecret string `json:"AccessSecret"` // JWT 密钥
	AccessExpire int64  `json:"AccessExpire"` // Token 过期时间（秒）
}

// PostgresqlConfig 定义 Postgresql 数据库配置
type PostgresqlConfig struct {
	DBHost     string `json:"db_host"`     // 数据库主机地址
	DBUser     string `json:"db_user"`     // 数据库用户名
	DBPassword string `json:"db_password"` // 数据库密码
	DBName     string `json:"db_name"`     // 数据库名称
	DBPort     string `json:"db_port"`     // 数据库端口
}
