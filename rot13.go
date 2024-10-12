package main

func rot13(input string) string {
	result := ""

	for _, c := range input {

		if c >= 'a' && c <= 'z' {
			result += string('a' + (c-'a'+13)%26)
		} else if c >= 'A' && c <= 'Z' {
			result += string('A' + (c-'A'+13)%26)
		} else {
			result += string(c)
		}
	}

	return result

}
