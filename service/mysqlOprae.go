package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"test/model"
)

func ShowMaster(master * model.Mysql)  {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", master.Username, master.Password, master.Host ,master.Dbname)

	db, err := sql.Open("mysql",url)
	if err != nil {
		fmt.Println(err)
	}
	r, _ := db.Query("show master status")

   	fmt.Println(r)
	//关闭数据库
	defer db.Close()
}
