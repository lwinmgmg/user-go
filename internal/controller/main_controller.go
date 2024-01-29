package controller

import (
	"time"

	"github.com/lwinmgmg/user-go/internal/services"
	redisctrl "github.com/lwinmgmg/user-go/pkg/redis-ctrl"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Controller struct {
	Db        *gorm.DB
	RoDb      *gorm.DB
	RedisCtrl *redisctrl.RedisCtrl
	LoginMail *services.MailService
	Otp       *services.OtpService
}

func NewContoller(db, roDb *gorm.DB, rd *redis.Client) *Controller {
	return &Controller{
		Db:        db,
		RoDb:      roDb,
		RedisCtrl: redisctrl.NewRedisCtrl(rd, time.Second*5),
	}
}
