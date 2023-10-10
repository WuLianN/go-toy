package global

import (
	"gorm.io/gorm"
	redis "github.com/redis/go-redis/v9"
)

var (
	DBEngine *gorm.DB
	RedisDBEngine *redis.Client
)
