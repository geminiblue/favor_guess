package middleware

import (
	"fmt"
	"github.com/geminiblue/favor_guess/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"log"
	"time"
)

var Db *xorm.Engine
var DbSlaver *xorm.Engine

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func InitDbConnection() {
	InitMaster()
	InitSlaver()
	go mysqlPingMaster()
	go mysqlPingSlaver()
}
func InitMaster() {
	var err error
	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.AppConfig.Db.User, config.AppConfig.Db.Password, config.AppConfig.Db.Host, config.AppConfig.Db.Port, config.AppConfig.Db.DataBase, config.AppConfig.Db.Charset)
	Db, err = xorm.NewEngine("mysql", connStr)
	FailOnError(err, "mysql 连接失败")
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "tbl_")
	Db.SetTableMapper(tbMapper)
	err = Db.RegisterSqlMap(xorm.Xml("./sql", ".xml"))
	FailOnError(err, "读取sql配置文件失败")
	Db.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	Db.ShowSQL(true)
	Db.ShowExecTime(true)
	Db.SetMaxIdleConns(config.AppConfig.Db.MaxIdleConns)
	Db.SetMaxOpenConns(config.AppConfig.Db.MaxOpenConns)
	err = Db.Ping()
	if err!=nil {
		log.Fatalf("数据主库连接失败%s",err.Error())
	}
}
func InitSlaver() {
	var err error
	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.AppConfig.DbSlaver.User, config.AppConfig.DbSlaver.Password, config.AppConfig.DbSlaver.Host, config.AppConfig.DbSlaver.Port, config.AppConfig.DbSlaver.DataBase, config.AppConfig.DbSlaver.Charset)
	DbSlaver, err = xorm.NewEngine("mysql", connStr)
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//DbSlaver.SetDefaultCacher(cacher)
	FailOnError(err, "mysql 连接失败")
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "tbl_")
	DbSlaver.SetTableMapper(tbMapper)
	err = DbSlaver.RegisterSqlMap(xorm.Xml("./sql", ".xml"))
	FailOnError(err, "读取sql配置文件失败")
	DbSlaver.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	DbSlaver.ShowSQL(true)
	DbSlaver.Logger().SetLevel(core.LOG_DEBUG)
	DbSlaver.SetMaxIdleConns(config.AppConfig.DbSlaver.MaxIdleConns)
	DbSlaver.SetMaxOpenConns(config.AppConfig.DbSlaver.MaxOpenConns)
	err = DbSlaver.Ping()
	if err!=nil {
		log.Fatalf("数据从库连接失败%s",err.Error())
	}
}
func mysqlPingMaster() {
	ticker := time.NewTicker(time.Second * time.Duration(60))
	for range ticker.C {
		err := Db.Ping()
		if err != nil {
			log.Printf("数据库连接丢失，原因%s", err.Error())
		}
	}
}
func mysqlPingSlaver() {
	ticker := time.NewTicker(time.Second * time.Duration(60))
	for range ticker.C {
		err := DbSlaver.Ping()
		if err != nil {
			log.Printf("数据库连接丢失，原因%s", err.Error())
		}
	}
}
