package iteration

// Repeat a given char 5 times
func Repeat(character string) (repeated string) {
	for counter := 0; counter < 5; counter++ {
		repeated += character
	}
	return
}
