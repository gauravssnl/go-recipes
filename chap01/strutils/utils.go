// Writing a Package as a Shared Library

package strutils

import (
	"strings"
	"unicode"
)

// functions starting with UppperCase act as public / are exportable
// function starting with LowerCase act as private / are not exported

// ToUpperCase : Convert string to uppercase
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// ToUpperCase : Convert string to lowercase
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// FirstToUpper : Convert first letter to uppercase
func FirstToupper(s string) string {
	if len(s) < 1 {
		return s
	}
	// removes space
	t := strings.Trim(s, " ")
	t = strings.ToLower(t)
	result := []rune(t) // alias for int32, distinguish character from int
	result[0] = unicode.ToUpper(result[0])
	return string(result)
}
