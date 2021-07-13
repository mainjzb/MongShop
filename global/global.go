package global

import (
	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB    *gorm.DB
	GVA_REDIS *redis.Client
	//GVA_CONFIG config.Server
	GVA_VP *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control = &singleflight.Group{}
)
