package stringutil

func Underscore(s string) string {
	var new string = ""
	for i := 0; i < len(s); i++ {
		new += string(s[i]) + "_"
	}
	new = new[:(len(s)-1)]
	return new
}
