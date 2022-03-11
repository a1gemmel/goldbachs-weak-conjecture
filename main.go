package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func ceil(x float64) int {
	return int(math.Ceil(x))
}

func sqrt(x int) float64 {
	return math.Sqrt(float64(x))
}

func primesBelow(n int) []int {
	notPrime := make([]bool, n)

	for i := 0; i < n; i++ {

		for divisor := 2; divisor <= min(i-1, ceil(math.Sqrt(float64(n)))); divisor++ {
			if i%divisor == 0 {
				notPrime[i] = true
				break
			}
		}
	}

	primes := []int{}

	for i, isntPrime := range notPrime {
		if !isntPrime && i > 1 {
			primes = append(primes, i)
		}
	}
	return primes
}

func goldbach(num int) [][]int {
	primes := primesBelow(num)

	solutions := [][]int{}

	for xi, x := range primes {

		for yi, y := range primes[xi:] {
			if x+y > num {
				break
			}
			for _, z := range primes[yi+xi:] {
				sum := x + y + z
				if sum == num {
					solutions = append(solutions, []int{x, y, z})
				}
				if sum >= num {
					break
				}
			}
		}
	}
	return solutions
}

func main() {

	num := flag.Int("sum", 0, "Number to sum to")
	flag.Parse()

	if *num == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *num%2 == 0 || *num < 7 {
		fmt.Println("Provided sum must be odd and greater than or equal to 7")
		os.Exit(1)
	}

	solutions := goldbach(*num)

	for _, s := range solutions {
		fmt.Println(s)
	}
}
