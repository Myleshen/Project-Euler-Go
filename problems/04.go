package problems

import (
	"sort"
)

func Problem4() int {
	operand1 := 999
	operand2 := 999
	var product_slice []int

	for {
		num := operand1 * operand2
		if IsPalindrome(num) {
			product_slice = append(product_slice, operand1*operand2)
			if operand1 > 100 {
				operand1--
			} else {
				operand1 = 999
				operand2--
			}
			if operand2 <= 100 {
				break
			}
		} else {
			if operand1 > 100 {
				operand1--
			} else {
				operand1 = 999
				operand2--
			}
			if operand2 <= 100 {
				break
			}
		}
	}
	sort.Slice(product_slice, func(i, j int) bool {
		return product_slice[i] > product_slice[j]
	})

	return product_slice[0]
}

func IsPalindrome(num int) bool {
	orig_num := num
	rev_num := 0
	for {
		if num > 0 {
			remainder := num % 10
			rev_num = (rev_num * 10) + remainder
			num = num / 10
		} else {
			break
		}
	}
	return rev_num == orig_num
}
