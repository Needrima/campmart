package helpers

func incrementPageNumber(num int) int {
	num++
	return num
}

func decrementPageNumber(num int) int {
	if num == 0 {
		return num
	}

	num--
	return num
}
