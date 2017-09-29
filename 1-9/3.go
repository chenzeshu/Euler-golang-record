//求600851475143的最大质因数
package main

import "fmt"

func getBiggestFactor(n int) int {
	factor := 2
	for n > 1{

		if(n % factor == 0){
			for n % factor == 0 {
				n /= factor
			}
		}
		if( factor*factor > 600851475143){
			return factor
		}
		if( n < factor){
			break
		}
		//如果factor的平方大于n ,那么n只有n能整除
		if( factor == 2){
			factor ++
		}else {
			factor += 2
		}
	}
	return factor
}

func main() {
	mySum := 600851475143
	var i int
	i = getBiggestFactor(mySum)

	fmt.Println(i)
}
