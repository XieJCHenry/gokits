package value

func GetOr[T any](cond bool, a T, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}
