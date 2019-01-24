package main

import (

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main(){
	//连接数据库
	 db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/productorder?charset=utf8")
	 //另两种方式打开数据库
	//db, err := sql.Open("mysql", "root:123@/productorder?charset=utf8")
	//	db, err := sql.Open("mysql", "root:123@/productorder")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	////执行更新数据语句
	//res,err:= db.Exec("update product set pprice = 2000 where pid = 9 ")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//fmt.Println(res.RowsAffected())
	//
	////执行删除语句
	//re,err := db.Exec("delete from product where pid = 9")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//fmt.Println(re.RowsAffected())
	//
	////执行插入语句
	//result,err := db.Exec("insert into product(pid,pname,pprice)values(9,'huawei',1600)")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//fmt.Println(result.RowsAffected())

	//执行数据查询语句
	//rows, err := db.Query ("select * from product where pid = 9") // ? = placeholder
	//rows, err := db.Query ("select * ,avg(pprice) from product group  by pid desc having pprice>3000 ") // ? = placeholder
	rows, err := db.Query ("select * ,avg(pprice) from product where pprice>3000 group  by pid desc  limit 1,4") // ? = placeholder

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close() // Close the statement when we leave main() / the program terminates

}
