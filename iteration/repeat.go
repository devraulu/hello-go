package iteration

import "strings"

func Repeat(text string) string {
	var repeated strings.Builder

	for i := 1; i <= 5; i++ {
		repeated.WriteString(text)
	}

	return repeated.String()
}
