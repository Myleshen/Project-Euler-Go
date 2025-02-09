package problems

func Problem5() int {
	start_num := 20
	for {
		if IsEvenlyDivisibile(start_num) {
			return start_num
		} else {
			start_num += 20
			if start_num >= 670442572800 {
				break
			}
		}
	}
	return 0
}

func IsEvenlyDivisibile(numerator int) bool {
	nums_to_divide_by := []int{20, 19, 17, 16, 15, 13, 11, 9, 8, 7, 6}
	for _, num := range nums_to_divide_by {
		if numerator%num != 0 {
			return false
		}
	}
	return true
}
