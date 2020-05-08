package iteration

import "strings"

// Repeat a given char n times
func Repeat(character string, numberOfRepeats int) (repeated string) {
	repeated = strings.Repeat(character, numberOfRepeats)
	return
}
