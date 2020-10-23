package main

import (
	"math/rand"
	"sync"
)

// NumberStorage for random number storage
type NumberStorage struct {
	numbers []int
	lock    sync.Mutex
}

// HighestOccuringNumbers returns the highest occuring numbers
//
// only returns multiple numbers if the amount of times they occurred are the same
func (nm *NumberStorage) HighestOccuringNumbers() []int {
	numberMap := make(map[int]int)
	highestNumbers := make(map[int]int)
	lastHighItem := 0

	nm.lock.Lock()
	defer nm.lock.Unlock()

	for _, num := range nm.numbers {
		if val, ok := numberMap[num]; ok {
			numberMap[num] = val + 1
		} else {
			numberMap[num] = 1
		}
		if len(highestNumbers) == 0 {
			highestNumbers[num] = numberMap[num]
			lastHighItem = num
		}
		// If the number is already in the map make a new map
		// Since its now the highest
		if _, ok := highestNumbers[num]; ok {
			highestNumbers = map[int]int{num: numberMap[num]}
			lastHighItem = num
		} else if numberMap[num] == highestNumbers[lastHighItem] {
			// Lets add the entry to the number map since it tied
			highestNumbers[num] = numberMap[num]
			lastHighItem = num
		}
	}

	numbers := make([]int, 0, len(highestNumbers))
	for k := range highestNumbers {
		numbers = append(numbers, k)
	}

	return numbers
}

func (nm *NumberStorage) worker(wg *sync.WaitGroup, worker int, batchCount int) {
	defer wg.Done()

	//fmt.Printf("Worker Start: %d\n", worker)

	numbers := GetRandomNumbers(worker, batchCount)

	nm.lock.Lock()
	nm.numbers = append(nm.numbers, numbers...)
	nm.lock.Unlock()

	//fmt.Printf("Worker Finshed: %d\n", worker)
}

// GetRandomNumbers a function that gets some random numbers
func GetRandomNumbers(worker int, batchCount int) []int {

	randNumbers := make([]int, 0, 1000)

	num := batchCount * worker
	rand.Seed(int64(num))

	times := num / (rand.Intn(num) + 1)

	for i := 0; i <= times; i++ {
		randNumbers = append(randNumbers, (rand.Intn(randomNumberMax) + randomNumberMin))
	}

	return randNumbers
}
