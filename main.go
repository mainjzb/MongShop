package main

import (
	"github.com/go-playground/validator/v10"
	"mongShop/config"
	"mongShop/global"
	"mongShop/router"
)

func main() {
	global.GVA_VP = config.Viper("./config/config.yaml")
	global.GVA_Validate = validator.New()
	global.GVA_DB = config.Gorm()
	router.Routers()

}
