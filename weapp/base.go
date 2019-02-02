package weapp

import (
	"github.com/medivhzhan/weapp"
	"log"
)

type MicroApp struct {
	AppID string
	AppSecret string
}


func (m MicroApp) GetToken()  {

}
//获取小程序登录验证信息，并保存session_key到redis
func (m MicroApp) GetAuthInfo(code string) (map[string]string,error) {
	res,err := weapp.Login(m.AppID,m.AppSecret,code)
	result := make(map[string]string)
	if err!=nil {
		log.Printf("获取用户验证信息失败,%s",err.Error())
		return nil,err
	}
	result["open_id"] = res.OpenID
	result["session_key"] = res.SessionKey
	result["UnionID"] = res.UnionID
	return result,nil
}
