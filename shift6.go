package main

func shift6(input string) string {
	result := ""

	for _, c := range input {

		if c >= 'a' && c <= 'z' {
			result += string('a' + (c-'a'+6)%26)
		} else if c >= 'A' && c <= 'Z' {
			result += string('A' + (c-'A'+6)%26)
		} else {
			result += string(c)
		}
	}

	return result

}
