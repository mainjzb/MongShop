package config

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	//Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	//Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	//Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	//Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	//// auto
	//AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
	//// gorm
	//Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//// oss
	//Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	//Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	//AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	//TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
	//Excel      Excel      `mapstructure:"excel" json:"excel" yaml:"excel"`
	//Timer      Timer      `mapstructure:"timer" json:"timer" yaml:"timer"`
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 端口值
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
}
