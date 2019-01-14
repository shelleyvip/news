package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"news/models"
	"path"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.html"
}

func (c *MainController) Post() {
	//1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//2.对数据进行校验
	if userName=="" || pwd == ""{
		beego.Info("数据不能为空")
		c.Redirect("/reg",302)
		return
	}
	//3.插入数据库
	o :=orm.NewOrm()

	user := models.User{}
	user.Name = userName
	user.Pwd = pwd

	_,err := o.Insert(&user) //一定是地址
	if err != nil{
		beego.Info("插入数据失败")
		c.Redirect("/reg",302)
		return
	}
	//c.Ctx.WriteString("插入成功")
	c.Redirect("/login",302)
}

/**
登录的get方法
 */
func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

/**
登录的Post方法  业务逻辑处理
 */
func (c *MainController) HandleLogin() {
	//拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	//判断数据是否合法
	if userName == "" || pwd == ""{
		beego.Info("输入的数据不合法")
		c.TplName = "login.html"
		return
	}
	//3.查询账户和密码是否正确

	o :=orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	err :=  o.Read(&user,"Name","Pwd") //select * from name=? and Pwd=?
	if err != nil{
		beego.Info("查询失败")
		c.Redirect("/login",302)
		return
	}
	// c.Ctx.WriteString("登录成功")
	c.Redirect("/index",302)


}

//显示列表功能
func (c *MainController) ShowIndex()  {
	//orm 查询
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []models.Article
	_,err :=qs.All(&articles)
	if err != nil{
		beego.Info("查询所有文章出错")
		return

	}
	count,err := qs.Count()
	if err != nil{
		beego.Info("查询错误")
		return
	}




	c.Data["count"] = count
	c.Data["articles"] =articles
	c.TplName = "index.html"
}

//文章添加get:showAdd;post:HandleAdd
func (c *MainController) ShowAdd() {
	c.TplName = "add.html"
}

func (c *MainController) HandleAdd(){
	//1.拿到数据
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")

	//文件上传功能
	f,h,err := c.GetFile("uploadname")
	defer f.Close()

	//1.限定格式 png jpg
	fileext := path.Ext(h.Filename) //取出后缀
	beego.Info(fileext)
		 	if fileext != ".img" && fileext != ".png" && fileext != ".jpeg" {
		beego.Info("上传文件格式错误")
		return
	}

	//2.限制大小
	if h.Size > 40000000{
		beego.Info("上传文件过大")
		return
	}

	//3.对文件重新命名，防止重复
	filename := time.Now().Format("2006-01-02") + fileext //6-1-2 3:4:5

	if err != nil{
		beego.Info("上传文件失败")
		fmt.Println("getfile err",err)
	}else {
		c.SaveToFile("uploadname","./static/img/"+filename)
	}

	if artiName == "" || artiContent == ""{
		beego.Info("添加文章数据错误")
		return
	}

	//3.插入数据库
	o := orm.NewOrm()
	arti := models.Article{}
	arti.ArtiName = artiName
	arti.Acontent = artiContent
	arti.Aimg = "/static/img/"+filename
	//c.Ctx.WriteString("添加文章成功")

	_,err = o.Insert(&arti)
	if err != nil{
		beego.Info("插入数据失败")
		return
	}
	//c.Ctx.WriteString("添加文章成功")
	c.Redirect("/index",302)


}

//显示内容详情页面
func (c *MainController) ShowContent()  {
    //1.获取文章ID
    id,err := c.GetInt("id")
    if err != nil{
    	beego.Info("获取文章Id错误",err)
    	return
	}
    //2.查询数据库对应的数据
    o:= orm.NewOrm()

    arti := models.Article{Id:id}
    err = o.Read(&arti)
    if err != nil{
    	beego.Info("查询错误",err)
    	return
	}
    //3.传递数据给视图
    c.Data["article"] = arti
    c.TplName = "content.html"

}

//显示页面编辑
func (c *MainController)ShowUpdate()  {
	//1.获取文章ID
	id,err :=c.GetInt("id")
	if err != nil{
		beego.Info("获取文章ID错误",err)
		return
	}
	//2.查询数据库功能
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询错误",err)
		return
	}
	//传递数据给视图
	c.Data["article"] =arti
	c.TplName = "update.html"


}

func (c *MainController) HandleUpdate()  {
	id,_ := c.GetInt("id")
	artiname:=c.GetString("articleName")
	content := c.GetString("content")
	f,h,err := c.GetFile("uploadname")

	defer f.Close()

	//1.限定格式 png jpg
	fileext := path.Ext(h.Filename) //取出后缀
	beego.Info(fileext)
	if fileext != ".jpg" && fileext != ".png" &&  fileext != ".jpeg"  {
		//if fileext != ".jpg" && fileext != ".png" &&  fileext != ".jpeg"  {

			beego.Info("上传文件格式错误")
		return
	}

	//2.限制大小
	if h.Size > 40000000{
		beego.Info("上传文件过大")
		return
	}

	//3.对文件重新命名，防止重复
	filename := time.Now().Format("2006-01-02") + fileext //6-1-2 3:4:5

	if err != nil{
		beego.Info("上传文件失败")
		fmt.Println("getfile err",err)
	}else {
		c.SaveToFile("uploadname","./static/img/"+filename)
	}
	//对数据进行一个处理
	if artiname == "" || content == "" {
		beego.Info("更新数据获取失败")
		return
	}

	//3.更新数据
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err  != nil{
		beego.Info("查询数据失败")
		return
	}
	arti.ArtiName = artiname
	arti.Acontent= content
	arti.Aimg = "./static/img/"+ filename

	_,err = o.Update(&arti,"ArtiName","Acontent","Aimg")
	if err != nil{
		beego.Info("更新数据显示错误")
		return
	}
	c.Redirect("/index",302)

}
//删除操作
func (c *MainController)HandelDelete(){
	//拿到数据
	id,err := c.GetInt("id")
	if err != nil{
		beego.Info("获取id数据错误")
		return
	}
	//执行删除操作
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("查询失败")
		return
	}
	o.Delete(&arti)
	c.Redirect("/index",302)

}



