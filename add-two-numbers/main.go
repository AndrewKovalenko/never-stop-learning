package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sumResult(num1 int, num2 int, overflow bool) (int, bool) {
	overflowAdjustment := 0

	if overflow {
		overflowAdjustment = 1
	}

	sum := num1 + num2 + overflowAdjustment

	if sum >= 10 {
		return sum % 10, true
	}

	return sum, false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	overflow := false

	var result *ListNode

	for {
		newNode := new(ListNode)

		num1 := 0
		num2 := 0

		if l1 != nil {
			num1 = l1.Val
		}

		if l2 != nil {
			num2 = l2.Val
		}

		var resultToSave int
		resultToSave, overflow = sumResult(num1, num2, overflow)
		newNode.Val = resultToSave

		if result == nil {
			result = newNode
		} else {
			newNode.Next = result
			result = newNode
		}

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			break
		}
	}

	return result
}

func main() {

}
