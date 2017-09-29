//求能被1-20每一个数字整除的最小数字
package main

import "fmt"

func main()  {
	//dividedNumber := [10]int{20, 19, 18, 17, 16, 15, 14, 13 ,12 ,11}
	dividedNumber := []int{20, 19, 18, 17, 16, 15, 14, 13 ,12 ,11}  //切片比较方便
	total := len(dividedNumber)
	found := false
	for i:= 2520; found == false; i++{
		found = true
		for j:= 0; j< total ; j++{
			if(i % dividedNumber[j] != 0){
				found = false
				break
			}
		}
		if(found){
			fmt.Println(i)
			break
		}
	}
}


