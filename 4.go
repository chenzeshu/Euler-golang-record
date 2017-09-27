//求三位数相乘能得到的最大回文数
package main

import "fmt"

//todo 判断是否为回文数
func hw(n int) bool{
	//获取n的反转数, 判断和n是否相等, 而不是去直接判断每一位是否相等(需要string, 而int也不允许)
	var _n int = 0
	var _injectN int = n
	for n != 0 {
		_n = _n * 10 + n%10
		n /= 10
	}
	if(_n == _injectN){
		return true
	}
	return false
}

func main()  {
	var sum, _sum int
	sum = 0
	for m:= 100; m<=999; m++{
		for n:= 100; n<= 999; n++{
			_sum = n * m
			if(hw(_sum)){
				if(sum < _sum){
					sum = _sum
				}
			}

		}
	}
	fmt.Println(sum)
}
