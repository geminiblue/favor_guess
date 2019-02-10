package config

import (
	"encoding/json"
	"fmt"
	"github.com/geminiblue/favor_guess/libs"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var AppConfig *Config
var Env string
var Secret string
var ApiAddr string

type Config struct {
	Application string         `json:"application"`
	Db          DbConfig       `json:"db"`
	DbSlaver    DbConfig       `json:"db_slaver"`
	MicroApp    MicroAppConfig `json:"micro_app"`
	Redis       cluster        `json:"redis"`
}
type DbConfig struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DataBase     string `json:"database"`
	User         string `json:"user"`
	Password     string `json:"password"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	Charset      string `json:"charset"`
}

//redis配置
type RedisConfig struct {
	Host                string `json:"host"`
	Port                string `json:"port"`
	Auth                string `json:"password"`
	MaxActiveConnection int    `json:"max_active_connection"`
}

//redis主从配置
type cluster struct {
	Master RedisConfig `json:"master"`
	Slaver RedisConfig `json:"slaver"`
}

//小程序配置
type MicroAppConfig struct {
	AppID         string `json:"app_id"`
	AppSecret     string `json:"app_secret"`
	Token         string `json:"token"`
	MchID         string `json:"mch_id"`
	MchAPISignKey string `json:"mch_api_sign_key"`
}

type RemoteConfig struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	Env       string    `json:"env"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
type RemoteResult struct {
	Data    *RemoteConfig `json:"data"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

//如果存在错误，则输出
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func GetConfig() *Config {
	url := fmt.Sprintf("https://geminiblue:gemini4094@c.oliyo.com/config/get/favor_guess/%s", Env)
	res, err := http.Get(url)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Fatalf("读取配置文件失败,[%s]文件不存在,%s", url, err.Error())
		//panic(err)
	}
	defer res.Body.Close()
	var remote RemoteResult
	body, err := ioutil.ReadAll(res.Body)
	failOnError(err, "读取配置文件失败,文件是空的")
	err = json.Unmarshal(body, &remote)
	failOnError(err, "读取配置文件失败,json解析失败")
	////encrypt_body := libs.Encrypt(string(body),key)
	////fmt.Println(encrypt_body)
	decryptBody := []byte(libs.Decrypt(string(remote.Data.Content), Secret))
	err = json.Unmarshal(decryptBody, &AppConfig)
	failOnError(err, "读取配置文件失败,json解析失败")
	return AppConfig
}
