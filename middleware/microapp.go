package middleware

import (
	"github.com/geminiblue/favor_guess/config"
	"github.com/geminiblue/favor_guess/weapp"
	"sync"
)

var once sync.Once
var microApp *weapp.MicroApp

//获取小程序通讯实例
func InsMicroApp() *weapp.MicroApp  {
	once.Do(func() {
		microApp = &weapp.MicroApp{
			AppID:config.AppConfig.MicroApp.AppID,
			AppSecret:config.AppConfig.MicroApp.AppSecret,
		}
	})
	return microApp
}