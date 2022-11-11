//你好，蓝山
/*
package main

import "fmt"

func main() {
	fmt.Println("Hello, Lanshan!")
}
*/
//变量
/*
var 变量名 类型 = 表达式

var a = "initial" // 类型推导，不指定类型自动判断

var b, c int = 1, 2 // 一次初始化多个变量

var d = true

var e float64 // 普通声明未赋值

f := float32(e) // 短声明

g := a + "apple"
fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
fmt.Println(g)                // initialapple
*/
//常量
/*
const s string = "constant"
const h = 500000000
const i = 3e20 / h
fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
*/
//练习619
/*
package main

import "fmt"

func main(){
	var L int
	fmt.Scanf("%d",&L)
	fmt.Printf("%v minutos", L*2)
}
*/
//for
/*
package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("loop")
		break // 跳出循环
	}

	// 打印7、8
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
			// 当n模2为0时不打印，进到下一次的循环
		}
		fmt.Println(n)
	}
	// 直到i>3
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
  // for 循环嵌套
  for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
*/
//练习723
/*
package main

import "fmt"

func main() {
    for{
        var X int
        fmt.Scanf("%d", &X)
        if X == 0{
            return
        }
        for i := 1; i <=X; i++{
            fmt.Printf("%d ", i)
        }
        fmt.Println()
    }
}
*/
//if
/*
if 条件表达式 {

	//当条件表达式结果为true时，执行此处代码
}

if 条件表达式 {
    //当条件表达式结果为true时，执行此处代码
} else {
    //当条件表达式结果为false时，执行此处代码
}
package main

import "fmt"

func main() {
	// 条件表达式为false，打印出"7 is odd"
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 条件表达式为ture，打印出"8 is divisible by 4"
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 短声明，效果等效于
	//num := 9
	//if num < 0{
	//	...
	//}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
*/
//练习664
/*package main

import "fmt"

func main() {
	var X, Y float32
	fmt.Scanf("%v%v", &X, &Y)
	if X > 0 && Y > 0 {
		fmt.Printf("Q1")
	} else if X < 0 && Y > 0 {
		fmt.Printf("Q2")
	} else if X < 0 && Y < 0 {
		fmt.Printf("Q3")
	} else if X > 0 && Y < 0 {
		fmt.Printf("Q4")
	} else if X == 0 && Y != 0 {
		fmt.Printf("Eixo Y")
	} else if X != 0 && Y == 0 {
		fmt.Printf("Eixo X")
	} else {
		fmt.Printf("Origem")
	}
}
*/
//switch
/*
package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		// 在此打印"two"并跳出
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	// 根据现在的时间判断是上午还是下午
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
*/
//练习661
/*
package main

import "fmt"

func main() {
	var i float32
	fmt.Scanf("%v", &i)
	switch {
	case i >= 0 && i <= 25:
		fmt.Printf("Intervalo [0,25]")
	case i > 25 && i <= 50:
		fmt.Printf("Intervalo (25,50]")
	case i > 50 && i <= 75:
		fmt.Printf("Intervalo (50,75]")
	case i > 75 && i <= 100:
		fmt.Printf("Intervalo (75,100]")
	default:
		fmt.Printf("Fora de intervalo")
	}
}
*/
//数组
/*
package main

import "fmt"

func main() {
	// 声明了长度为5的数组，数组中的每一个元素都是int类型
	var a [5]int
	// 给数组a的第4位元素赋值为100
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	// 在给数组声明的同时赋值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 声明二位数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
*/
//练习739
/*
package main

import "fmt"

func main() {
	var X [10]int
	var i, a int
	for i = 0; i < 10; i++ {
		fmt.Scanf("%v", &a)
		if a <= 0 {
			a = 1
		}
		X[i] = a
		fmt.Printf("X[%v] = %v\n", i, X[i])
	}
}
*/
//练习743
/*
package main

import "fmt"

	func main() {
		var N, T, i, r int
		var X [60]int
		X[0] = 0
		X[1] = 1
		for i = 2; i < 60; i++ {
			X[i] = X[i-1] + X[i-2]
		}
		fmt.Scanf("%v", &T)
		for r = 0; r < T; r++ {
			fmt.Scanf("%v", &N)
			fmt.Printf("Fib(%v) = %v\n", N, X[N])
		}
	}
*/
//slice
