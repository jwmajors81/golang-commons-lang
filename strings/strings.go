package utils

import (
	"errors"
	"math"
	"regexp"
	"strings"
	"unicode"

	"github.com/jwmajors81/golang-commons-lang/sorted"
)

const (
	LF = "\n"
	CR = "\r"
)

// Abbreviates a String using ellipses.  This will turn
// "It's a dangerous business, Frodo, going out your door" into
// "It's a dangerous business..."
func Abbreviate(value string, maxWidth int) (string, error) {
	if maxWidth >= len(value) {
		return value, nil
	} else if maxWidth < 4 {
		return "", errors.New("the maxWidth must be greater than 3")
	} else if (len(value) - 3) > 0 {
		runes := []rune(value)
		return string(runes[:maxWidth-3]) + "...", nil
	} else if (len(value) - 3) == 0 {
		runes := []rune(value)
		return string(runes[:maxWidth-3]), nil
	}

	return value, nil
}

// Appends the suffix to the end of the string if the string does not already
// end with the suffix
func AppendIfMissing(value string, suffix string, ignoreCase bool) string {
	if ignoreCase {
		if HasSuffixIgnoreCase(value, suffix) {
			return value
		}
	} else {
		if strings.HasSuffix(value, suffix) {
			return value
		}
	}

	return value + suffix

}

// Returns true if the string already ends with the specified suffix (case-insensitive)
func HasSuffixIgnoreCase(value string, suffix string) bool {
	valueLower := strings.ToLower(value)
	suffixLower := strings.ToLower(suffix)

	return strings.HasSuffix(valueLower, suffixLower)
}

// Capitalizes a string by changing the first character to title case
func Capitalize(value string) string {
	firstLetter := Substr(value, 0, 1)
	restOfValue := SubstrRight(value, 1)

	return strings.ToUpper(firstLetter) + restOfValue
}

// Returns the end of the string starting with the index specified
func SubstrRight(input string, start int) string {
	return Substr(input, start, len(input)-start)
}

// Returns the beginning of the string ending with the index specified
func SubstrLeft(input string, end int) string {
	if end <= 0 {
		return ""
	}

	end = sorted.Min(len(input), end)
	return Substr(input, 0, end)
}

// Returns the substring of the string based upon the start position and length required
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start > len(input) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// Removes one new line from end of string if it's there.  New lines are considered:
// \n, \r, or \r\n
func RemoveLastSeparator(original string) string {
	length := len(original)
	if length == 0 {
		return ""
	}

	lastCharacter := Substr(original, length-1, 1)
	if length > 1 && lastCharacter == LF && Substr(original, length-2, 1) == CR {
		return Substr(original, 0, length-2)
	} else if lastCharacter == LF {
		return Substr(original, 0, length-1)
	} else if lastCharacter == CR {
		return Substr(original, 0, length-1)
	}

	return original
}

// Remove the last character from a string
func Chop(original string) string {
	length := len(original)
	if length < 1 {
		return ""
	}

	lastCharacter := Substr(original, length-1, 1)
	if length > 1 && lastCharacter == LF && Substr(original, length-2, 1) == CR {
		return Substr(original, 0, length-2)
	}

	return Substr(original, 0, len(original)-1)
}

// Centers a string in a larger string
func Center(original string, size int, char rune) (string, error) {
	if size < 0 {
		return original, nil
	}

	padding := size - len(original)
	if padding <= 0 {
		return original, nil
	}

	var err error
	var ret string
	if ret, err = LeftPad(original, len(original)+padding/2, char); err != nil {
		return "", err
	}

	if ret, err = RightPad(ret, size, char); err != nil {
		return "", err
	}

	return ret, nil

}

// Add padding to the left of the original string using the value specified
func LeftPad(original string, size int, char rune) (string, error) {
	if !unicode.IsPrint(char) {
		return "", errors.New("the padding character must be a printable character as defined by unicode.IsPrint")
	}

	charNum := size - len(original)
	if charNum < 0 {
		return original, nil
	}

	leftPad := strings.Repeat(string(char), charNum)
	return leftPad + original, nil
}

// Add padding to the right of the string using the value specified
func RightPad(original string, size int, char rune) (string, error) {
	if !unicode.IsPrint(char) {
		return "", errors.New("the padding character must be a printable character as defined by unicode.IsPrint")
	}

	charNum := size - len(original)
	if charNum < 0 {
		return original, nil
	}

	rightPad := strings.Repeat(string(char), charNum)
	return original + rightPad, nil
}

// Checks whether particular strings aren't found in a string
func ContainsNone(original string, searchFor ...string) bool {
	for _, val := range searchFor {
		if len(val) == 0 {
			// do nothing and process next value
		} else if strings.Contains(original, val) {
			return false
		}
	}
	return true
}

// Checks whether the only values found in a string are those that are specified
func ContainsOnly(original string, searchFor ...string) bool {
	stringToSearch := original
	for _, val := range searchFor {
		stringToSearch = strings.ReplaceAll(stringToSearch, val, "")
	}

	return len(stringToSearch) == 0
}

// Finds the first non-empty string and returns the value if found, otherwise nil is returned
func FirstNonEmpty(values ...string) *string {
	for _, val := range values {
		if len(val) > 0 {
			return &val
		}
	}

	return nil
}

func SafeDeref(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

// Compares all strings and returns the initial sequence of chracters that are common
// to all of the strings
func CommonPrefix(values ...string) string {
	if len(values) == 0 {
		return ""
	} else if len(values) == 1 {
		return values[0]
	}

	firstValue := values[0]
	lastPrefix := ""

	for _, char := range firstValue {
		currentPrefix := lastPrefix + string(char)

		for _, value := range values {
			if !strings.HasPrefix(value, currentPrefix) {
				return lastPrefix
			}
		}

		lastPrefix = currentPrefix
	}

	return lastPrefix
}

// Returns all digits found in the string
func GetDigits(original string) string {
	r := regexp.MustCompile("[0-9]")
	matches := r.FindAllString(original, -1)
	if matches == nil {
		return ""
	}

	return strings.Join(matches, "")
}

// Search input for the first matching value found and returns the index of that value
func IndexOfAny(original string, searchFor ...string) int {
	for _, val := range searchFor {
		if index := strings.Index(original, val); index > 0 {
			return index
		}
	}
	return -1
}

// Compares all values specified and returns the index where the values start to differ
func IndexOfDifference(values ...string) int {
	if len(values) == 0 {
		return -1
	}

	valueLengths := LengthOfStrings(values...)
	shortestVal := sorted.Min(valueLengths...)
	longestVal := sorted.Max(valueLengths...)

	firstValue := values[0]

	for i := 0; i < shortestVal; i++ {
		compareTo := firstValue[i]
		for _, val := range values {
			if val[i] != compareTo {
				return i
			}
		}
	}

	if longestVal > shortestVal {
		return shortestVal
	}

	return -1
}

// Retruns the length of all of the strings provided
func LengthOfStrings(values ...string) []int {
	var lengthOfValues []int
	for _, val := range values {
		lengthOfValues = append(lengthOfValues, len(val))
	}

	return lengthOfValues
}

// Finds the last index that matches the string that is provided and -1 if no matches are found
func LastIndexOf(original string, searchFor string) int {
	return LastIndexOfWithStartPos(original, searchFor, sorted.Max(0, len(original)))
}

// Finds the last index that matches the striung that is provided after the start position.
// If no match is found then -1 is returned.
func LastIndexOfWithStartPos(original string, searchFor string, startPosition int) int {
	if startPosition > len(original) {
		startPosition = len(original)
	}
	if startPosition < 0 {
		return -1
	}
	original = original[0:startPosition]
	for charPos := startPosition; charPos >= 0; charPos-- {
		searchIn := original[charPos:]
		if index := strings.Index(searchIn, searchFor); index >= 0 {
			return charPos
		}
	}

	return -1
}

// Finds the latest index of any of the values specified within the original string
func LastIndexOfAny(original string, searchFor ...string) int {
	var results []int
	for _, val := range searchFor {
		results = append(results, LastIndexOf(original, val))
	}

	return sorted.Max(results...)
}

// Gets the rightmost 'x' characters of a string
func Right(original string, length int) string {
	if original == "" {
		return ""
	}

	if length > len(original) {
		return original
	}

	return original[len(original)-length:]
}

// Rotate (circular shift) a string 'x' times based upon the input.
func Rotate(original string, shift int) string {
	if len(original) == 0 {
		return original
	}

	shiftAbs := math.Abs(float64(shift))

	if int(shiftAbs) > len(original) {
		shift = shift % len(original)
	}

	chars := strings.Split(original, "")
	shifted := make([]string, len(chars))

	for index, char := range chars {
		newPosition := index + shift

		if newPosition >= 0 && newPosition < len(chars) {
			shifted[newPosition] = char
		} else if newPosition >= len(chars) {
			correctedPosition := newPosition - len(chars)
			shifted[correctedPosition] = char
		} else if newPosition < 0 {
			correctedPosition := len(chars) + newPosition
			shifted[correctedPosition] = char
		}
	}

	return strings.Join(shifted, "")
}

// Determines whether one string starts within another string
func StartsWith(searchFor string, searchIn string, ignoreCase bool) bool {
	if len(searchFor) == 0 {
		return false
	}

	if len(searchFor) > len(searchIn) {
		return false
	}

	if ignoreCase {
		searchFor = strings.ToLower(searchFor)
		searchIn = strings.ToLower(searchIn)
	}

	for i := 0; i < len(searchFor); i++ {
		lookingFor := searchFor[i]
		actual := searchIn[i]

		if lookingFor != actual {
			return false
		}
	}
	return true
}

// Returns the substring of the original string after the first instance of the separate is found.
// Returns an empty string if the value cannot be found.
func SubstrAfter(original string, separator rune) string {
	index := strings.IndexRune(original, separator)

	if index < 0 {
		return ""
	}

	return SubstrRight(original, index+1)
}

// Returns the substring of the original string after the last instance of the separate is found.
// Returns an empty string if the value cannot be found.
func SubstrAfterLast(original string, separator rune) string {

	index := strings.LastIndexFunc(original, func(compareTo rune) bool {
		return compareTo == separator
	})

	if index < 0 {
		return ""
	}

	return SubstrRight(original, index+1)
}

// Returns the substring of the original string before the first instance of the separate is found.
// Returns an empty string if the value cannot be found.
func SubstrBefore(original string, separator rune) string {
	index := strings.IndexRune(original, separator)

	if index < 0 {
		return ""
	}

	return SubstrLeft(original, index)
}
