package app

import (
	"log"

	"github.com/geminiblue/favor_guess/api"
	"github.com/geminiblue/favor_guess/config"
	"github.com/geminiblue/favor_guess/middleware"
)

type App struct {
	Addr   string
	Env    string
	Secret string
}

//初始化配置
func (app App) setConfig() {
	config.Env = app.Env
	config.Secret = app.Secret
	config.ApiAddr = app.Addr
	config.GetConfig()
	log.Printf("%s 程序开始执行", config.AppConfig.Application)
}

//初始化系列连接池
func (app App) initMiddleware() {
	middleware.InitDbConnection()
	middleware.InitRedisConnection()
}

//程序运行
func (app App) Run() {
	app.setConfig()
	app.initMiddleware()
	go api.Run()
	select {}
}
