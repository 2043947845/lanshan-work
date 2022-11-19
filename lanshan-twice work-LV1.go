package main

import (
	"fmt"
	"time"
)

type 好康的 func() string //先用"好康的"代表下面两个相同类型的函数

func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() string {
	return "输了啦，都是你害的"
}

func Activity(fType 好康的) (result string) { //fType表示一个函数类型变量
	result = fType()
	return
}
func main() {
	fmt.Println(Activity(打电动))
	fmt.Println(Activity(欢迎来我家玩))
}
