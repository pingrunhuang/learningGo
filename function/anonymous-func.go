package function

// below defined a squares function that return a function with return value as integer
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
