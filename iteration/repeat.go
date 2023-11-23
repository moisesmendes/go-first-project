package iteration

func Repeat(character string, n int) string {
	var repeated string // declaring a variable
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}
