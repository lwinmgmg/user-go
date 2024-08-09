package controller

import (
	"time"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/services"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	otpctrl "github.com/lwinmgmg/user-go/pkg/otp-ctrl"
	redisctrl "github.com/lwinmgmg/user-go/pkg/redis-ctrl"
	"gorm.io/gorm"
)

type Controller struct {
	Db           *gorm.DB
	RoDb         *gorm.DB
	RedisCtrl    *redisctrl.RedisCtrl
	LoginMail    *services.MailService
	PhoneService *services.PhoneService
	Otp          *services.OtpService
	JwtCtrl      *jwtctrl.JwtCtrl
	Setting      *env.Settings
}

func NewContoller(settings *env.Settings) *Controller {
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		panic(err)
	}
	roDb, err := services.GetPsql(settings.RoDb)
	if err != nil {
		panic(err)
	}
	rd, err := services.GetRedisClient(settings.Redis)
	if err != nil {
		panic(err)
	}
	return &Controller{
		Db:           db,
		RoDb:         roDb,
		RedisCtrl:    redisctrl.NewRedisCtrl(rd, time.Second*5),
		LoginMail:    services.NewMailService(settings.LoginEmailServer),
		PhoneService: services.NewPhoneServer(),
		Otp: services.NewOtpService(&otpctrl.OtpCtrl{
			Issuer: settings.Service,
		}, otpctrl.STANDARD_OPT_DURATION, settings.OtpService.Skew),
		JwtCtrl: jwtctrl.NewJwtCtrl(settings.Service),
		Setting: settings,
	}
}
