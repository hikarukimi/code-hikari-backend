package svc

import (
	"code-hikari/common-go"
	model "code-hikari/common-go/model"
	"code-hikari/user/rpc/internal/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	Redis     *redis.Client
	TokenUtil *common.UserClaims
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 返回初始化的 ServiceContext
	return &ServiceContext{
		Config:    c,
		DB:        initPostgres(c),
		Redis:     initRedis(c),
		TokenUtil: initTokenUtil(c),
	}
}
func initPostgres(c config.Config) *gorm.DB {
	// 构建 Postgresql 数据库连接字符串
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.PostgresqlConfig.DBHost, c.PostgresqlConfig.DBUser, c.
			PostgresqlConfig.DBPassword, c.PostgresqlConfig.DBName, c.PostgresqlConfig.DBPort)

	// 使用 GORM 连接 Postgresql
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// 测试数据库连接是否成功
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(err.Error())
	}

	// 自动迁移模型
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&model.UserRelation{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&model.UserSignIn{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&model.UserFavorite{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the Postgresql And Successfully auto migrate models!")
	return db
}

func initRedis(c config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	fmt.Println("Successfully connected to the Redis!")
	return rdb
}

func initTokenUtil(c config.Config) *common.UserClaims {
	return common.NewTokenUtil(c.TokenAuth.AccessSecret, c.TokenAuth.AccessExpire)
}
