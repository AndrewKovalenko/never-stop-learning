package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num1length := len(nums1)
	num2Length := len(nums2)

	mergedArrayLength := num1length + num2Length
	resultHasEvenLength := mergedArrayLength%2 == 0

	medianElementIndex := mergedArrayLength / 2

	resultingArray := make([]int, medianElementIndex+1)
	num1Index := 0
	num2Index := 0

	for i := 0; i <= medianElementIndex; i++ {
		if num1Index > num1length-1 {
			resultingArray[i] = nums2[num2Index]
			num2Index++
			continue
		}

		if num2Index > num2Length-1 {
			resultingArray[i] = nums1[num1Index]
			num1Index++
			continue
		}

		if nums1[num1Index] < nums2[num2Index] {
			resultingArray[i] = nums1[num1Index]
			num1Index++
		} else {
			resultingArray[i] = nums2[num2Index]
			num2Index++
		}
	}

	resultingArrayLength := len(resultingArray)

	if resultHasEvenLength {
		lastElement := resultingArray[resultingArrayLength-1]
		beforeLastElement := resultingArray[resultingArrayLength-2]

		return float64(lastElement+beforeLastElement) / 2
	} else {
		return float64(resultingArray[resultingArrayLength-1])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{2, 3}))
}
