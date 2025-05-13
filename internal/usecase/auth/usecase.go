package authucase

import (
	"glossika/internal/service/redis"

	"gorm.io/gorm"
)

type Usecase struct {
	redisSvc *redis.RedisService
	db       *gorm.DB
}

func New(redisSvc *redis.RedisService, db *gorm.DB) *Usecase {
	return &Usecase{
		redisSvc: redisSvc,
		db:       db,
	}
}
