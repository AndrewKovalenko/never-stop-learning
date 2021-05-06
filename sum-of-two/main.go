package main

import "fmt"

var secondArgumentIndexes = make(map[int]int)

func twoSum(nums []int, target int) []int {
	for index, item := range nums {
		secondArgumentValue := target - item

		if secondArgIndex, exists := secondArgumentIndexes[secondArgumentValue]; exists && index != secondArgIndex {
			result := [...]int{secondArgIndex, index}

			return result[:]
		}

		secondArgumentIndexes[item] = index
	}

	return nil
}

func main() {
	nums := [...]int{-10, 7, 19, 15}
	target := 9
	fmt.Println(twoSum(nums[:], target))
}
