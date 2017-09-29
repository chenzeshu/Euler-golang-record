//勾股数
package main

import (
	"fmt"
)

func main() {
	//两处优化:
	//1: a与b的最大值都不会超过500
	//2: 因为a,b都是自然数, 且a<b<c, 所以 b = a+1
	//运算次数为79576次
	var a, b, c ,times int
	times = 0;
	for a= 1; a< 500; a++{
		for b=(a+1); b< 500 ; b++{
			c = 1000 - a - b
			times++
			if( a*a +b*b == c*c){
				fmt.Println(a*b*c,times)
				return
			}
		}
	}

}