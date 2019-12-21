//没做出来，正则表达式不会，思路就是用正则表达式找出学号和姓名，
// 然后匹配学号尾数就行，名字俩字就是检查名字中间的空格数
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

)
var Message struct{
	XH int
	XM string
}
func main() {
	//var s2 []Message
	data, err := ioutil.ReadFile("E:/Students.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println(string(data))
	chiReg := regexp.MustCompile("[^\u4e00-\u9fa5]")
	s1 := chiReg.ReplaceAllString()
	fmt.Println(s1)
	reg := regexp.MustCompile((?<=\").*(?=\"))
	//s1:=reg.FindAllString(str, -1)
	//s2=append(s2,s1)
	fmt.Println("------FindAll------")
	//返回str中所有匹配reg的字符串
	//第二个参数表示最多返回的个数，传-1表示返回所有结果
	dataSlice := reg.FindAll([]byte(str), -1)
	for _, v := range dataSlice {
		fmt.Println(string(v))
	}
}

//`(xh":").* ","xb`
	//pattern := `xh":">(.*?)<","xm":">`
	//rp2 := regexp.MustCompile(pattern)
	//fmt.Println(rp2)


