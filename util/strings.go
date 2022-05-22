package util

func RepeatString(s string, count int) string {
	r := ""
	for i := 0; i < count; i++ {
		r += s
	}
	return r
}
