package iteration

func Repeat(char string, times int) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += char
	}

	return repeated
}
