package module01

// check if bla bla
func NumInList(list []int, num int) bool{
	for _, i := range list {
		if i == num {
		return true
	}
}
	return false
}