package main

import "fmt"

func main (){
	i, j:= 0, 3
	prime := []int{2}

	for i< 10001 {
		tmp := j
		for _, v :=range prime {
			if( tmp % v == 0){
				tmp = 0
				break
			}
		}
		if ( tmp != 0){
			prime = append(prime, j)
			i++
		}
		j++
	}

	fmt.Println(prime[10000])
}
