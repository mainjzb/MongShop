package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GVA_DB    *gorm.DB
	GVA_REDIS *redis.Client

	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Validate            *validator.Validate
	GVA_Concurrency_Control = &singleflight.Group{}
	//GVA_LOG    *oplogging.Logger
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
)
