package main

import "fmt"

func main()  {
	sum1, sum2:=0, 0
	for i,j:=1,1 ; i<=100 ; i++{
		sum1 += i*i
		sum2 += j
		if(j == 100){
			sum2 = sum2 * sum2
			break
		}
		j++
	}

	fmt.Println(sum2  - sum1)
}
