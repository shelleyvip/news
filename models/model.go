package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Name string
	Pwd string
	//Age string

	//Article[]*Article `orm:"rel(m2m)"`//设置多对多
}

type Article struct {
	Id int `orm:"pk;auto"` //主键 并且自动增长
	ArtiName string `orm:"size(20)"` //ArtiName长度为20
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string
	Aimg string
	// n 记住1的一方
	ArticleType *ArticleType `orm:"rel(fk)"` //设置一对多的关系

	//User[]*User `orm:"reverse(many)"` //设置多对多反向关系

}
//类型表
type ArticleType struct {
	Id int
	Tname string
	Article []*Article `orm:"reverse(many)"`  //设置一对多的反向关系
}


func init()  {
	//设置数据库的基本信息
	orm.RegisterDataBase("default","mysql","root:123@tcp(127.0.0.1:3306)/test5?charset=utf8")
	//映射modle数据
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	//生成表
	orm.RunSyncdb("default",false,true)
}