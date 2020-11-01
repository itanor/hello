package iteration

func Repeat(times int, term string) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += term
	}
	return
}
