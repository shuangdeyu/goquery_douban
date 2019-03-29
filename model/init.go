package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
)

const (
	DbUser     = "root"
	DbPassword = "root"
	DbName     = "goquery_douban"
	DbPreFix   = "q_"
)

var Engine *xorm.Engine

type Arr map[string]interface{}

func DbInit() *xorm.Engine {
	if Engine == nil {
		// 连接数据库
		engine, err := xorm.NewEngine("mysql", DbUser+":"+DbPassword+"@/"+DbName+"?charset=utf8")
		if err != nil {
			log.Println("数据库连接出错: ", err.Error())
		}
		// 设置表前缀映射
		if DbPreFix != "" {
			tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, DbPreFix)
			engine.SetTableMapper(tbMapper)
		}
		Engine = engine
	}
	return Engine
}
