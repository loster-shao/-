package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

func main() {
	var s1 string
	var s2 int
	var food []string
	var money [] int
	food = []string{"小面", "饭团", "香菇滑鸡", "小炒", "黄焖鸡", "冒菜"}
	money = []int{6, 7, 12, 15, 16, 18}
	for i := 0; i < len(food); i++ {
		s1=food[i]
		s2=money[i]
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/texting?charset=utf8")
		if err != nil {
			panic(err)
		}
		stmt, err := db.Prepare("INSERT foods SET food=?,money=? ")
		if err != nil {
			fmt.Println(err, "..")
		}
		_, err = stmt.Exec(s1,s2)
		if err != nil {
			fmt.Println("fail to insert",err)
		}
	}
	var a int
	rand.Seed(time.Now().UnixNano())
	 a =rand.Intn(5)
	//随机数范围 [0,n)
	fmt.Println("此次随机点餐为：",food[a])
}
