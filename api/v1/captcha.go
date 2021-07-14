package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"mongShop/config"
	"mongShop/global"
	"mongShop/model/response"
)

var store = base64Captcha.DefaultMemStore

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(config.GVA_CONFIG.Captcha.ImgHeight, config.GVA_CONFIG.Captcha.ImgWidth, config.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
