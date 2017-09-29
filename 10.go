package main

import (
	"fmt"
	"math"
)

func getPrime(n int64) []int64 {
	var result = make([]int64, 0, n/int64( math.Log(float64(n) )))
	var sieve = make([]bool, n+1)
	var sn = int64( math.Sqrt(float64(n)) )
	var i, j int64
	for i = 2; i<= sn; i++{
		if !sieve[i] {
			for j = i+i ; j<=n; j += i{
				sieve[j] = true
			}
		}
	}

	for i = 2; i<=n ; i++{
		if !sieve[i] {
			result = append(result, i)
		}
	}

	return result
}

func main()  {
	primes := getPrime(2e6)
	var sum int64 = 0
	for _, k:= range primes{
		sum += k
	}
	fmt.Println(sum)
}