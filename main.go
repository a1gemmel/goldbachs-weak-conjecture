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

	primes := primesBelow(*num)

	for x := 0; x < len(primes); x++ {
		xP := primes[x]
		for y := x; y < len(primes); y++ {
			yP := primes[y]
			if xP+yP > *num {
				break
			}
			for z := y; z < len(primes); z++ {
				zP := primes[z]
				sum := xP + yP + zP
				if sum == *num {
					fmt.Println(xP, yP, zP)
				}
				if sum >= *num {
					break
				}
			}
		}
	}
}
