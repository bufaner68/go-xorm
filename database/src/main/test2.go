package main

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"github.com/go-xorm/core"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Username string    `xorm:"not null VARCHAR(32)"`
	Birthday time.Time `xorm:"DATE"`
	Sex      string    `xorm:"CHAR(1)"`
	Address  string    `xorm:"VARCHAR(256)"`
}

func main(){
	//连接数据库
	engine,err := xorm.NewEngine("mysql","root:123@/productorder?charset=utf8")
	if err != nil{
		fmt.Println(err)
		return
	}
	//测试连通性
	err = engine.Ping()
	if err != nil{
		fmt.Println(err)
		return
	}
	//日志打印SQL
	engine.ShowSQL(true)

	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})

	////进行数据添加
	user := new(User)
	//user.Id = 3
	//user.Username = "tim"
	//user.Address = "shandong"
	//affected,err := engine.Insert(user)
	//fmt.Println(affected)
	//fmt.Println(user.Id)
	//fmt.Println("================")

	////进行数据查询1:sql语句
	//sql := "select * from user"
	//engine.Query(sql)
	//results,err := engine.Query(sql)
	//for i := range  results{
	//	fmt.Println(i,"值是：",results[i])
	//}
	//fmt.Println("SQL语句：******************************")

	////进行数据查询2：使用get（）方法：判断某条记录是否存在，若存在，则返回这条记录
	////根据ID获取数据,直接把查询语句和ID打印
	has,err := engine.Id(1).Get(user)
	fmt.Println(has)
	if has{
		fmt.Println(user)
	}

	fmt.Println("where:******************************")
	//根据where来获取，直接把查询语句和记录打印
	has,err = engine.Where("username=?","Amy").Get(user)
	fmt.Println(has)

	fmt.Println("exist:******************************")
	//使用exist（）方法：判断某条记录是否存在
	user = &User{Id:2}
	has,err = engine.Exist(user)
	fmt.Println(has)
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println("insert:******************************")
	//使用SQL语句进行数据的删除，修改，增加
	sql := "insert into user(id,username) values(4,'nina')"
	res,err := engine.Exec(sql)
	fmt.Println(res)
	if err != nil {
		fmt.Println("插入数据错误")
	}
}