package global

import (
	"CSAwork/model/config"
	"context"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config    *config.Config
	Logger    *zap.Logger
	GlobalDb1 *gorm.DB
	RedisDb   *redis.Client
	Ctx       context.Context
)
