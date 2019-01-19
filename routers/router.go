package routers

import (
	"news/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

)

func init() {
	//第一个参数: /article/* 路由规则，可以根据一定的规则进行路由，如果你全匹配可以用*
	//第二个参数:position执行Filter的地方 beego.BeforeRouter 寻找路由之前操作
	//第三个参数:执行的具体函数 beforExecFun
	beego.InsertFilter("/article/*", beego.BeforeRouter, beforExecFunc)//在执行列表上进行操作
	beego.Router("/", &controllers.MainController{})
	beego.Router("/reg", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/article/index", &controllers.MainController{}, "get:ShowIndex")
	beego.Router("/article/addArticle", &controllers.MainController{}, "get:ShowAdd;post:HandleAdd")
	beego.Router("/article/content", &controllers.MainController{}, "get:ShowContent")
	beego.Router("/article/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/article/delete", &controllers.MainController{}, "get:HandelDelete")
	beego.Router("/article/addType", &controllers.MainController{}, "get:ShowAddType;post:HandleAddType")
	beego.Router("/logout", &controllers.MainController{}, "get:LogOut")

	}
//过滤器的判断
        var beforExecFunc = func(ctx *context.Context) {
        	 userName := ctx.Input.Session("userName")
        	 if userName == nil{
        	 	ctx.Redirect(302,"/login")
			 }

		}