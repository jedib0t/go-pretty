package util

// FilterStrings filters the slice 's' to items which return truth when passed
// to 'f'.
func FilterStrings(s []string, f func(string) bool) []string {
	var out []string
	for _, item := range s {
		if f(item) {
			out = append(out, item)
		}
	}
	return out
}
