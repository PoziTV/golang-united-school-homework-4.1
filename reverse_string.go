package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	runeOutput := []rune(input)

	// len_2 := len(output) / 2

	for i := 0; i < len(runeOutput)/2; i++ {
		runeOutput[i], runeOutput[len(runeOutput)-1-i] = runeOutput[len(runeOutput)-1-i], runeOutput[i]
	}

	return string(runeOutput) //output
}
