package config

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"log"
	"mongShop/global"
	"mongShop/model"
	"os"
	"time"
)

func Gorm() *gorm.DB {
	switch GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
	//	model.SysAuthority{},
	//	model.SysApi{},
	//	model.SysBaseMenu{},
	//	model.SysBaseMenuParameter{},
	//	model.JwtBlacklist{},
	//	model.SysDictionary{},
	//	model.SysDictionaryDetail{},
	//	model.ExaFileUploadAndDownload{},
	//	model.ExaFile{},
	//	model.ExaFileChunk{},
	//	model.ExaSimpleUploader{},
	//	model.ExaCustomer{},
	//	model.SysOperationRecord{},
	//
	//	// Code generated by gin-vue-admin Begin; DO NOT EDIT.
	//	// Code generated by gin-vue-admin End; DO NOT EDIT.
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}

//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch GVA_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = Default.LogMode(logger.Info)
	default:
		config.Logger = Default.LogMode(logger.Info)
	}
	return config
}

type config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}

var (
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), config{})
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	Recorder = traceRecorder{Interface: Default, BeginAt: time.Now()}
)

func New(writer logger.Writer, config config) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s\n"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s\n"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s\n"
	)

	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s\n"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s\n" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s\n"
	}

	return &_logger{
		Writer:       writer,
		config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type _logger struct {
	config
	logger.Writer
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (c *_logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *c
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (c *_logger) Info(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Info {
		c.Printf(c.infoStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (c *_logger) Warn(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Warn {
		c.Printf(c.warnStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (c *_logger) Error(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Error {
		c.Printf(c.errStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (c *_logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if c.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && c.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				c.Printf(c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > c.SlowThreshold && c.SlowThreshold != 0 && c.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", c.SlowThreshold)
			if rows == -1 {
				c.Printf(c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case c.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				c.Printf(c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

func (c *_logger) Printf(message string, data ...interface{}) {
	if GVA_CONFIG.Mysql.LogZap {
		global.GVA_LOG.Info(fmt.Sprintf(message, data...))
	} else {
		c.Writer.Printf(message, data...)
	}
}

type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func (t traceRecorder) New() *traceRecorder {
	return &traceRecorder{Interface: t.Interface, BeginAt: time.Now()}
}

func (t *traceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	t.BeginAt = begin
	t.SQL, t.RowsAffected = fc()
	t.Err = err
}
