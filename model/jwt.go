package model

type JwtBlacklist struct {
	GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
