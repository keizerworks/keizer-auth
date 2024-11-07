package utils

import "unicode"

func ToSnakeCase(str string) string {
	var result []rune
	for i, char := range str {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(char))
	}
	return string(result)
}
