package reverse_string

func ReverseString(input string) (output string) {
	runes_in := []rune(input)
	var runes_out []rune
	for i := len(runes_in) - 1; i <= 0; i-- {
		runes_out = append(runes_out, runes_in[i])
	}
	output = string(runes_out)
	return output
}
