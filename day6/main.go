package main

import (
	"fmt"
	"math/big"
	. "math/big"
)

func countFish(buckets []*Int) *Int {
	count := big.NewInt(0)

	for _, b := range buckets {
		count = count.Add(count, b)
	}
	return count
}

func age(buckets []*Int) []*Int {
	Newbuckets := append(buckets[1:], buckets[0])
	Newbuckets[6] = Newbuckets[6].Add(Newbuckets[6], Newbuckets[len(Newbuckets)-1])
	return Newbuckets
}

func task1and2(fishes []int, days int) (result *Int) {
	buckets := make([]*Int, 9)
	for a := 0; a < len(buckets); a++ {
		buckets[a] = NewInt(0)
	}
	for _, f := range fishes {
		buckets[f] = buckets[f].Add(buckets[f], big.NewInt(1))
	}
	for day := 0; day < days; day++ {
		buckets = age(buckets)
	}
	return countFish(buckets)
}

func main() {
	input := "input.txt"

	fishes := readdata(input)
	days := 80
	fmt.Printf("Task 1 - # fishes after %d days   = %d \n", days, task1and2(fishes, days))
	days = 256
	fmt.Printf("Task 2 - # fishes after %d days  = %d \n", days, task1and2(fishes, days))
}
