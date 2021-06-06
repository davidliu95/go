package main

import (
	"database/sql"
	"github.com/pkg/errors"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//Db数据库连接池
var DB *sql.DB

type User struct {
	id   int64
	name string
	age  int8
}

func main() {
	err := InitMysql()
	defer DB.Close()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
	result, err := Query()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
	fmt.Println(result)

}

func InitMysql() error {
	DB, err := sql.Open("mysql", "studytest:Asd123456@tcp(rm-bp1g4908zi29h2rzwyo.mysql.rds.aliyuncs.com:3306)/study?charset=utf8")
	if err != nil {
		return errors.Wrap(err, "connnect sql err")
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(20)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(5)
	//验证连接
	if err := DB.Ping(); err != nil {
		return errors.Wrap(err, "db ping err")
	}
	fmt.Println("connnect success")
	return nil
}

//查询操作
func Query() (result []User, err error) {
	var users []User
	rows, e := DB.Query("select * from user where id in (1,2,3)")
	if e != nil {
		if e == sql.ErrNoRows {
			return users, nil
		}
		return users, errors.Wrap(e, "query users err")
	}
	for rows.Next() {
		var user User
		e := rows.Scan(user.name, user.id, user.age)
		if e != nil {
			return users, errors.Wrap(e, "query users scan err")
		}
		users = append(users, user)
	}
	err = rows.Close()
	if err != nil {
		return users, errors.Wrap(err, "query users close rows err")
	}
	return users, nil
}
