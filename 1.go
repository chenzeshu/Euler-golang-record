//100以内3和5的公倍数之和

package main

import "fmt"

func getSum() int {
	sum := 0;
	for i:=0; i<1000; i++{
		if( i % 15 == 0){
			sum += i
			continue
		}
		if( i% 5 == 0){
			sum += i
			continue
		}
		if( i% 3 == 0){
			sum += i
			continue
		}
	}
	return sum
}

func main() {
	var mySum int
	mySum = getSum()

	fmt.Println(mySum)
}