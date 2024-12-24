package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

func getRandomInt(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var nums []int

	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}

	for _, num := range nums {
		ch <- num
	}

	close(ch)
}

func getPower(ch chan int, wg *sync.WaitGroup, nums *[]int) {
	defer wg.Done()

	for num := range ch {
		*nums = append(*nums, int(math.Pow(float64(num), 2)))
	}
}

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup

	wg.Add(1)
	go getRandomInt(ch, &wg)

	var nums []int

	wg.Add(1)
	go getPower(ch, &wg, &nums)

	wg.Wait()

	var numsPowered []string

	for _, num := range nums {
		numsPowered = append(numsPowered, strconv.Itoa(num))
	}

	fmt.Println(strings.Join(numsPowered, " "))
}
