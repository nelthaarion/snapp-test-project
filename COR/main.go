package main

import "log"

func main() {
	numbers := []int{10, 5, 2}
	powered := Power(numbers...)
	counter := len(numbers)
	finalResult := 0
	for num := range Sum(powered) {
		counter--
		if counter == 0 {
			finalResult = num
		}
	}
	log.Println(finalResult)
}

func Power(nums ...int) <-chan int {
	powered := make(chan int)
	go func() {
		for _, n := range nums {
			powered <- n * n
		}
		close(powered)
	}()
	return powered
}
func Sum(powered <-chan int) <-chan int {
	result := make(chan int)
	sum := 0
	go func() {
		for n := range powered {
			sum += n
			result <- sum
		}
		close(result)
	}()
	return result
}
