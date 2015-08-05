package fibonacci

func IsFib(n int64) bool {
	if n == 0 {
		return true
	}

	var twoNumbersAgo int64 = 0
	var oneNumberAgo int64 = 0

	var i int64
	for i = 1; i < n; i = twoNumbersAgo + oneNumberAgo {
		twoNumbersAgo = oneNumberAgo
		oneNumberAgo = i
	}

	return i == n
}
