package help

func Half(num int) (int, bool) {
	return num / 2, (num%2 == 0)
}
