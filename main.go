package main

import (
	"flag"
	"github.com/geminiblue/favor_guess/app"
	"runtime"
)

var (
	env  string
	key  string
	addr string
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&env, "env", "dev", "启动环境")
	flag.StringVar(&key, "key", "gemini4094", "启动秘钥")
	flag.StringVar(&addr, "addr", "127.0.0.1:19840", "api服务启动端口")
	flag.Parse()
}
func main() {
	application := app.App{
		Env:    env,
		Secret: key,
		Addr:   addr,
	}

	application.Run()
}
