//斐波那契数400万内偶数项求和
package main

import "fmt"

func getSum() int{
	sum,f1,f2,f3 := 0,0,1,0
	for n:=0 ; f2 < 4e6 ; n++ {
		f3 = f1 + f2
		if( f3 % 2 ==0){
			sum += f3
		}
		f1 = f2
		f2 = f3
	}
	return sum
}


func main()  {
	var sum int
	sum = getSum()
	fmt.Println(sum)
}
