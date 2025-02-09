package problems

func Problem6(natural_num int) int {
	sum_of_squares := (natural_num * (natural_num + 1)) / 2
	sum_of_squares *= sum_of_squares
	square_of_sum := (natural_num * (natural_num + 1) * (2*natural_num + 1)) / 6
	if square_of_sum > sum_of_squares {
		return square_of_sum - sum_of_squares
	}
	return sum_of_squares - square_of_sum
}
