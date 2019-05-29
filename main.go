package main

import (
	_ "myRestProject3/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm" // 引入beego的orm
	_ "github.com/go-sql-driver/mysql" // 引入mysql driver
	// "fmt"
	// "github.com/gomodule/redigo/redis" // 引入第三方redis
	// "github.com/astaxie/beego/cache" // 引入redis
	// _ "github.com/astaxie/beego/cache/redis" // 引入redis
)

func main() {
	
	beego.SetLogger("file", `{"filename":"D:/isbprs_golang_log.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	
	// BConfig是beego框架的默认配置，对应app.conf的配置
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	orm.RegisterDataBase("default", "mysql",
        "root:Pass1234@tcp(10.203.103.31:3306)/isbprs?charset=utf8", 30)
	// orm.RegisterModel()  // 是在每个model里注册的
	// o := orm.NewOrm()  // 是在每个model里insert的
    // user1 := Test{Idtest: 111}
    // id, err1 := o.Insert(&user1)
    
    // redis connection
    // conn, _ := redis.Dial("tcp", "10.203.104.178:7527") //rongxing给我的服务器
    
    beego.Debug("start backend server")
    
	beego.Run()
}
