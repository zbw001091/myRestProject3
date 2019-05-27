package main

import (
	_ "myRestProject3/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm" // 引入orm
	_ "github.com/go-sql-driver/mysql" // 引入mysql driver
	"fmt"
)

func main() {
	
	// BConfig是beego框架的默認配置，對應app.conf的配置
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	orm.RegisterDataBase("default", "mysql",
        "root:Pass1234@tcp(10.203.103.31:3306)/isbprs?charset=utf8", 30)
	// orm.RegisterModel()  // 是在每個model里註冊的
	// o := orm.NewOrm()  // 是在每個model里insert的
    // user1 := Test{Idtest: 111}
    // id, err1 := o.Insert(&user1)
    
    fmt.Printf("start backend server")
    
	beego.Run()
}
