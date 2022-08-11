package main

import "log"

func main() {
	s := []int{3, 18, 3, 2, 2, 4, 4, 5, 6, 6, 5, 7, 8, 7, 8}

	log.Println(FindOrphan(s))
}

func FindOrphan(data []int) int {
	result := 0
	resultMap := map[int]int{}
	for _, v := range data {
		if resultMap[v] == 0 {
			resultMap[v] = v
			result += v
		} else {
			delete(resultMap, v)
			result -= v
		}

	}
	return result
}
