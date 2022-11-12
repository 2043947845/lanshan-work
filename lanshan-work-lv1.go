//两个版本
// 第一版，涉及阶乘，运算时间久
package main

import (
	"fmt"
)

// 阶乘
func j(w int) int {
	res := 1
	for i := 1; i <= w; i++ {
		res *= i
	}
	return res
}

func main() {
	var n, l, r, q, w int
	fmt.Scanf("%v %v %v", &n, &l, &r)
	a := make([]int, n)
	//给数组里n个值赋值
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &q)
		a[i] = q
	}
	if l <= r && r <= n {
		w = r - l                   //只排l,r之间的数
		for t := 0; t < j(w); t++ { //进行w！次从头到尾的排序
			for e := 0; e < w; e++ { //进行一次从头到尾的排序
				if a[l+e] > a[l+1+e] { //对前后数比较并排序
					a[l+e], a[l+1+e] = a[l+1+e], a[l+e]
				}
			}
		}
	}
	for m := 0; m < n; m++ {
		fmt.Printf("%v ", a[m])
	}
}

//第二版，冒泡排序
package main

import "fmt"

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	var n, l, r, q int
	fmt.Scanf("%v %v %v", &n, &l, &r)
	a := make([]int, n)
	b := make([]int, r-l+1)
	//给数组里n个值赋值
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &q)
		a[i] = q
	}
	for c := 0; c <= r-l; c++{//将a[r]~a[l]提出给b[]
		b[c] = a[l+c]
	}
	bubbleSort(b)//对b[]冒泡排序，只排a[r]~[l]
	for d := 0; d <= r-l; d++{//将排序好的a[r]~a[l]放回原数组
		a[l+d] = b[d]
	}
	for m := 0; m < n; m++ {
		fmt.Printf("%v ", a[m])
	}
}