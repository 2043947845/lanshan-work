package main //w！的时间复杂程度很大，我目前只能通过阶乘次将数组中的数一遍又一遍的排，竭尽全力了。

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
	var n, l, r, q, w, x int
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
					x = a[l+e]
					a[l+e] = a[l+1+e]
					a[l+1+e] = x
				}
			}
		}
	}
	for m := 0; m < n; m++ {
		fmt.Printf("%v ", a[m])
	}
}
