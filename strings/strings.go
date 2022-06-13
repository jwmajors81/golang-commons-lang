package utils

import (
	"errors"
	"math"
	"regexp"
	"strings"

	"github.com/jwmajors81/golang-commons-lang/sorted"
)

const (
	LF = "\n"
	CR = "\r"
)

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

func HasSuffixIgnoreCase(value string, suffix string) bool {
	valueLower := strings.ToLower(value)
	suffixLower := strings.ToLower(suffix)

	return strings.HasSuffix(valueLower, suffixLower)
}

func Capitalize(value string) string {
	firstLetter := Substr(value, 0, 1)
	restOfValue := SubstrLeft(value, 1)

	return strings.ToUpper(firstLetter) + restOfValue
}

func SubstrLeft(input string, start int) string {
	return Substr(input, start, len(input)-start)
}

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

func Chop(input string) string {
	length := len(input)
	if length < 2 {
		return ""
	}

	newLength := length - 1

	ret := Substr(input, 0, newLength)
	lastCharacter := Substr(input, newLength, 1)
	if lastCharacter == LF && SubstrLeft(ret, newLength-1) == CR {
		return Substr(ret, 0, newLength-1)
	}

	return ret
}

func Center(original string, size int, char string) (string, error) {
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

func LeftPad(original string, size int, char string) (string, error) {
	if len(char) != 1 {
		return "", errors.New("the padding character must have a length of 1")
	}

	charNum := size - len(original)
	if charNum < 0 {
		return original, nil
	}

	leftPad := strings.Repeat(char, charNum)
	return leftPad + original, nil
}

func RightPad(original string, size int, char string) (string, error) {
	if len(char) != 1 {
		return "", errors.New("the padding character must have a length of 1")
	}

	charNum := size - len(original)

	if charNum < 0 {
		return original, nil
	}

	rightPad := strings.Repeat(char, charNum)
	return original + rightPad, nil
}

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

func ContainsOnly(original string, searchFor ...string) bool {
	stringToSearch := original
	for _, val := range searchFor {
		stringToSearch = strings.ReplaceAll(stringToSearch, val, "")
	}

	return len(stringToSearch) == 0
}

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

func GetDigits(original string) string {
	r := regexp.MustCompile("[0-9]")
	matches := r.FindAllString(original, -1)
	if matches == nil {
		return ""
	}

	return strings.Join(matches, "")
}

func IndexOfAny(original string, searchFor ...string) int {
	for _, val := range searchFor {
		if index := strings.Index(original, val); index > 0 {
			return index
		}
	}
	return -1
}

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

func LengthOfStrings(values ...string) []int {
	var lengthOfValues []int
	for _, val := range values {
		lengthOfValues = append(lengthOfValues, len(val))
	}

	return lengthOfValues
}

func LastIndexOf(original string, searchFor string) int {
	return LastIndexOfWithStartPos(original, searchFor, sorted.Max(0, len(original)))
}

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

func LastIndexOfAny(original string, searchFor ...string) int {
	var results []int
	for _, val := range searchFor {
		results = append(results, LastIndexOf(original, val))
	}

	return sorted.Max(results...)
}

func Right(original string, length int) string {
	if original == "" {
		return ""
	}

	if length > len(original) {
		return original
	}

	return original[len(original)-length:]
}

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
