package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbbreviate(t *testing.T) {
	var tests = map[string]struct {
		input          string
		maxWidth       int
		expectedOutput string
		expectedError  error
	}{
		"width is greater than input": {input: "a", maxWidth: 2, expectedOutput: "a", expectedError: nil},
		"empty":                       {input: "", maxWidth: 2, expectedOutput: "", expectedError: nil},
		"width is less than length of string input": {input: "abcdefg", maxWidth: 6, expectedOutput: "abc...", expectedError: nil},
		"width is equals length of string input":    {input: "abcdefg", maxWidth: 7, expectedOutput: "abcdefg", expectedError: nil},
		"width is greater than string input 2":      {input: "abcdefg", maxWidth: 8, expectedOutput: "abcdefg", expectedError: nil},
		"invalid max width":                         {input: "abcdefg", maxWidth: 3, expectedOutput: "", expectedError: errors.New("the maxWidth must be greater than 3")},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Abbreviate(test.input, test.maxWidth)

			if test.expectedError != nil {
				assert.Equal(t, "", actual)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedOutput, actual)
			}
		})
	}
}

func TestCenter(t *testing.T) {
	var tests = map[string]struct {
		original       string
		size           int
		paddingChar    rune
		expectedOutput string
		expectedError  error
	}{
		"empty":                             {original: "", size: 4, paddingChar: ' ', expectedOutput: "    ", expectedError: nil},
		"negative size":                     {original: "ab", size: -1, paddingChar: ' ', expectedOutput: "ab", expectedError: nil},
		"size greater than length of input": {original: "ab", size: 4, paddingChar: ' ', expectedOutput: " ab ", expectedError: nil},
		"size less than length of input":    {original: "abcd", size: 2, paddingChar: ' ', expectedOutput: "abcd", expectedError: nil},
		"padding is space":                  {original: "a", size: 4, paddingChar: ' ', expectedOutput: " a  ", expectedError: nil},
		"padding is y":                      {original: "a", size: 4, paddingChar: 'y', expectedOutput: "yayy", expectedError: nil},
		"padding is yb":                     {original: "a", size: 4, paddingChar: '\n', expectedOutput: "", expectedError: errors.New("the padding character must be a printable character as defined by unicode.IsPrint")},
	}

	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Center(val.original, val.size, val.paddingChar)
			if val.expectedError == nil {
				assert.Equal(t, val.expectedOutput, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, val.expectedOutput, actual)
				assert.Equal(t, val.expectedError, err)
			}
		})
	}
}

func TestLeftPad(t *testing.T) {
	var tests = map[string]struct {
		original       string
		size           int
		paddingChar    rune
		expectedOutput string
		expectedError  error
	}{
		"empty":                                        {original: "", size: 3, paddingChar: ' ', expectedOutput: "   ", expectedError: nil},
		"negative size":                                {original: "ab", size: -1, paddingChar: ' ', expectedOutput: "ab", expectedError: nil},
		"size is same as length of input":              {original: "bat", size: 3, paddingChar: ' ', expectedOutput: "bat", expectedError: nil},
		"size is greater than length of input":         {original: "bat", size: 5, paddingChar: ' ', expectedOutput: "  bat", expectedError: nil},
		"size is less than length of input":            {original: "bat", size: 1, paddingChar: ' ', expectedOutput: "bat", expectedError: nil},
		"padding character's length is greater than 1": {original: "a", size: 4, paddingChar: '\n', expectedOutput: "", expectedError: errors.New("the padding character must be a printable character as defined by unicode.IsPrint")},
	}

	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := LeftPad(val.original, val.size, val.paddingChar)
			if val.expectedError == nil {
				assert.Equal(t, val.expectedOutput, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, val.expectedOutput, actual)
				assert.Equal(t, val.expectedError, err)
			}
		})
	}
}

func TestRightPad(t *testing.T) {
	var tests = map[string]struct {
		original       string
		size           int
		paddingChar    rune
		expectedOutput string
		expectedError  error
	}{
		"empty":                              {original: "", size: 3, paddingChar: ' ', expectedOutput: "   ", expectedError: nil},
		"negative size":                      {original: "ab", size: -1, paddingChar: ' ', expectedOutput: "ab", expectedError: nil},
		"same size":                          {original: "bat", size: 3, paddingChar: ' ', expectedOutput: "bat", expectedError: nil},
		"size greater than input":            {original: "bat", size: 5, paddingChar: ' ', expectedOutput: "bat  ", expectedError: nil},
		"size less than input":               {original: "bat", size: 1, paddingChar: ' ', expectedOutput: "bat", expectedError: nil},
		"padding greater than one character": {original: "a", size: 4, paddingChar: '\n', expectedOutput: "", expectedError: errors.New("the padding character must be a printable character as defined by unicode.IsPrint")},
	}

	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := RightPad(val.original, val.size, val.paddingChar)
			if val.expectedError == nil {
				assert.Equal(t, val.expectedOutput, actual)
				assert.Nil(t, err)
			} else {
				assert.Equal(t, val.expectedOutput, actual)
				assert.Equal(t, val.expectedError, err)
			}
		})
	}
}

func TestAppendIfMissing(t *testing.T) {
	var tests = map[string]struct {
		input          string
		suffix         string
		ignoreCase     bool
		expectedOutput string
	}{
		"a -> ab":                     {input: "a", suffix: "b", ignoreCase: true, expectedOutput: "ab"},
		"empty input and suffix":      {input: "", suffix: "", ignoreCase: true, expectedOutput: ""},
		"empty suffix":                {input: "abc", suffix: "", ignoreCase: true, expectedOutput: "abc"},
		"empty input, but not suffix": {input: "", suffix: "abc", ignoreCase: true, expectedOutput: "abc"},
		"abc + xyz -> abcxyz":         {input: "abc", suffix: "xyz", ignoreCase: true, expectedOutput: "abcxyz"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expectedOutput, AppendIfMissing(test.input, test.suffix, test.ignoreCase))
		})
	}
}

func TestHasSuffixIgnoreCase(t *testing.T) {
	assert.True(t, HasSuffixIgnoreCase("abc", "BC"))
	assert.True(t, HasSuffixIgnoreCase("abc", "bc"))
	assert.False(t, HasSuffixIgnoreCase("abc", "AB"))
	assert.False(t, HasSuffixIgnoreCase("abc", "ab"))
	assert.True(t, HasSuffixIgnoreCase("a ", " "))
	assert.True(t, HasSuffixIgnoreCase("", ""))
}

func TestCapitalize(t *testing.T) {
	assert.Equal(t, Capitalize(""), "")
	assert.Equal(t, Capitalize("cat"), "Cat")
	assert.Equal(t, Capitalize("cAt"), "CAt")
	assert.Equal(t, Capitalize("'cat'"), "'cat'")
}

func TestSubstrLeft(t *testing.T) {
	assert.Equal(t, "", SubstrLeft("abc", 0))
	assert.Equal(t, "a", SubstrLeft("abc", 1))
	assert.Equal(t, "abc", SubstrLeft("abc", 5))
	assert.Equal(t, "", SubstrLeft("", 2))
}

func TestSubstrRight(t *testing.T) {
	assert.Equal(t, "abc", SubstrRight("abc", 0))
	assert.Equal(t, "bc", SubstrRight("abc", 1))
	assert.Equal(t, "", SubstrRight("abc", 5))
	assert.Equal(t, "", SubstrRight("", 2))
}

func TestSubstrBefore(t *testing.T) {
	assert.Equal(t, "", SubstrBefore("", 'a'))
	assert.Equal(t, "", SubstrBefore("abc", 'a'))
	assert.Equal(t, "ab", SubstrBefore("abcba", 'c'))
	assert.Equal(t, "ab", SubstrBefore("abc", 'c'))
	assert.Equal(t, "", SubstrBefore("abcba", 'a'))
	assert.Equal(t, "", SubstrBefore("abcba", 'd'))
}

func TestSubstrAfter(t *testing.T) {
	assert.Equal(t, "", SubstrAfter("", 'a'))
	assert.Equal(t, "bc", SubstrAfter("abc", 'a'))
	assert.Equal(t, "ba", SubstrAfter("abcba", 'c'))
	assert.Equal(t, "", SubstrAfter("abc", 'c'))
	assert.Equal(t, "", SubstrAfter("abcba", 'd'))
}

func TestSubstrAfterLast(t *testing.T) {
	assert.Equal(t, "", SubstrAfterLast("", 'a'))
	assert.Equal(t, "bc", SubstrAfterLast("abc", 'a'))
	assert.Equal(t, "a", SubstrAfterLast("abcba", 'b'))
	assert.Equal(t, "a", SubstrAfterLast("abcbacba", 'b'))
	assert.Equal(t, "", SubstrAfterLast("abc", 'c'))
	assert.Equal(t, "", SubstrAfterLast("abcba", 'd'))
}

func TestSubstringWithLength(t *testing.T) {
	assert.Equal(t, "a", Substr("abc", 0, 1))
	assert.Equal(t, "abc", Substr("abc", 0, 3))
	assert.Equal(t, "abc", Substr("abc", 0, 5))
	assert.Equal(t, "", Substr("", 1, 2))
}

func TestChop(t *testing.T) {
	assert.Equal(t, "", Chop(""))
	assert.Equal(t, "", Chop("a"))
	assert.Equal(t, "abc ", Chop("abc \n"))
	assert.Equal(t, "abc ", Chop("abc \n"))
	assert.Equal(t, "abc ", Chop("abc \r\n"))
	assert.Equal(t, "ab", Chop("abc"))
	assert.Equal(t, "abc\nab", Chop("abc\nabc"))
	assert.Equal(t, "", Chop("a"))
	assert.Equal(t, "", Chop("\r"))
	assert.Equal(t, "", Chop("\n"))
	assert.Equal(t, "", Chop("\r\n"))

}

func TestRemoveLastSeparator(t *testing.T) {
	assert.Equal(t, "", RemoveLastSeparator(""))
	assert.Equal(t, "abc ", RemoveLastSeparator("abc "))
	assert.Equal(t, "abc ", RemoveLastSeparator("abc \r"))
	assert.Equal(t, "abc ", RemoveLastSeparator("abc \n"))
	assert.Equal(t, "abc", RemoveLastSeparator("abc\r\n"))
	assert.Equal(t, "abc\n", RemoveLastSeparator("abc\n\r"))
	assert.Equal(t, "abc", RemoveLastSeparator("abc"))
	assert.Equal(t, "abc\nabc", RemoveLastSeparator("abc\nabc"))
	assert.Equal(t, "abc\nabc", RemoveLastSeparator("abc\nabc\n"))
	assert.Equal(t, "a", RemoveLastSeparator("a"))
	assert.Equal(t, "", RemoveLastSeparator("\r"))
	assert.Equal(t, "", RemoveLastSeparator("\n"))
	assert.Equal(t, "", RemoveLastSeparator("\r\n"))
}

func TestContainsNone(t *testing.T) {
	assert.True(t, ContainsNone("abc", ""))
	assert.True(t, ContainsNone("abc", " "))
	assert.True(t, ContainsNone("", ""))
	assert.False(t, ContainsNone("abc", "abc"))
	assert.False(t, ContainsNone("abc", "c"))
}

func TestContainsOnly(t *testing.T) {
	assert.False(t, ContainsOnly("aaa", "b"))
	assert.True(t, ContainsOnly("aaa", "a"))
	assert.False(t, ContainsOnly("abad", "a", "b"))
	assert.True(t, ContainsOnly("abad", "a", "b", "d"))
	assert.False(t, ContainsOnly("a", ""))
}

func TestFirstNonEmpty(t *testing.T) {
	assert.Equal(t, "a", *FirstNonEmpty("", "a", "b"))
	assert.Nil(t, FirstNonEmpty("", "", ""))
	assert.Equal(t, "a", *FirstNonEmpty("a", "b"))
	assert.Equal(t, "a", *FirstNonEmpty("a", ""))
}

func TestSafeDeref(t *testing.T) {
	assert.Equal(t, "", SafeDeref(nil))
	value := "foo"
	assert.Equal(t, "foo", SafeDeref(&value))
}

func TestCommonPrefix(t *testing.T) {
	assert.Equal(t, "", CommonPrefix(""))
	assert.Equal(t, "", CommonPrefix("", ""))
	assert.Equal(t, "abc", CommonPrefix("abc"))
	assert.Equal(t, "", CommonPrefix("abc", ""))
	assert.Equal(t, "", CommonPrefix("", "", "abc"))
	assert.Equal(t, "abc", CommonPrefix("abc", "abc"))
	assert.Equal(t, "a", CommonPrefix("abc", "a"))
	assert.Equal(t, "ab", CommonPrefix("ab", "abefg"))
	assert.Equal(t, "ab", CommonPrefix("abcd", "abefg"))
	assert.Equal(t, "", CommonPrefix("abcd", "xyz"))
	assert.Equal(t, "", CommonPrefix("xyz", "abcde"))
	assert.Equal(t, "i am a ", CommonPrefix("i am a machine", "i am a robot"))
}

func TestGetDigits(t *testing.T) {
	assert.Equal(t, "", GetDigits(""))
	assert.Equal(t, "", GetDigits("abc"))
	assert.Equal(t, "1000", GetDigits("1000$"))
	assert.Equal(t, "112345", GetDigits("1123~45"))
	assert.Equal(t, "5417543010", GetDigits("(541) 754-3010"))
}

func TestIndexOfAny(t *testing.T) {
	assert.Equal(t, -1, IndexOfAny("abc", ""))
	assert.Equal(t, 2, IndexOfAny("zzabyy", "ab", "cd"))
	assert.Equal(t, 2, IndexOfAny("zzabyy", "cd", "ab"))
	assert.Equal(t, -1, IndexOfAny("zzabyy", "mn", "zz"))
	assert.Equal(t, -1, IndexOfAny("", "a"))
}

func TestIndexOfDifference(t *testing.T) {
	assert.Equal(t, -1, IndexOfDifference("abc"))
	assert.Equal(t, -1, IndexOfDifference("", ""))
	assert.Equal(t, 0, IndexOfDifference("abc", "", ""))
	assert.Equal(t, 0, IndexOfDifference("", "", "abc"))
	assert.Equal(t, -1, IndexOfDifference("abc", "abc"))
	assert.Equal(t, 1, IndexOfDifference("a", "abc"))
	assert.Equal(t, 2, IndexOfDifference("ab", "abxyz"))
	assert.Equal(t, 2, IndexOfDifference("abcde", "abxyz"))
	assert.Equal(t, 0, IndexOfDifference("abcde", "xyz"))
	assert.Equal(t, 0, IndexOfDifference("xyz", "abcde"))
	assert.Equal(t, 7, IndexOfDifference("i am a machine", "i am a robot"))
}

func TestLastIndexOf(t *testing.T) {
	assert.Equal(t, 0, LastIndexOf("", ""))
	assert.Equal(t, 7, LastIndexOf("aabaabaa", "a"))
	assert.Equal(t, 5, LastIndexOf("aabaabaa", "b"))
	assert.Equal(t, 4, LastIndexOf("aabaabaa", "ab"))
	assert.Equal(t, 8, LastIndexOf("aabaabaa", ""))
}

func TestLastIndexOfWithStartPos(t *testing.T) {
	assert.Equal(t, -1, LastIndexOfWithStartPos("", "", -1))
	assert.Equal(t, 5, LastIndexOfWithStartPos("aabaabaa", "b", 8))
	assert.Equal(t, 2, LastIndexOfWithStartPos("aabaabaaz", "b", 4))
	assert.Equal(t, -1, LastIndexOfWithStartPos("aabaabaa", "b", 0))
	assert.Equal(t, 5, LastIndexOfWithStartPos("aabaabaa", "b", 9))
	assert.Equal(t, -1, LastIndexOfWithStartPos("aabaabaa", "b", -1))
	assert.Equal(t, 0, LastIndexOfWithStartPos("aabaabaa", "a", 1))
}

func TestLastIndexOfAny(t *testing.T) {
	assert.Equal(t, 6, LastIndexOfAny("zzabyycdxx", "ab", "cd"))
	assert.Equal(t, 6, LastIndexOfAny("zzabyycdxx", "cd", "ab"))
	assert.Equal(t, -1, LastIndexOfAny("zzabyycdxx", "mn", "op"))
	assert.Equal(t, 10, LastIndexOfAny("zzabyycdxx", "mn", ""))
}

func TestRight(t *testing.T) {
	assert.Equal(t, "", Right("", 2))
	assert.Equal(t, "", Right("abc", 0))
	assert.Equal(t, "bc", Right("abc", 2))
	assert.Equal(t, "abc", Right("abc", 4))
}

func TestRotate(t *testing.T) {
	assert.Equal(t, "cab", Rotate("abc", 1))
	assert.Equal(t, "bca", Rotate("abc", 2))
	assert.Equal(t, "abc", Rotate("abc", 3))
	assert.Equal(t, "cab", Rotate("abc", 4))
	assert.Equal(t, "abc", Rotate("abc", 0))
	assert.Equal(t, "bca", Rotate("abc", -1))
	assert.Equal(t, "cab", Rotate("abc", -2))
	assert.Equal(t, "abc", Rotate("abc", -3))
	assert.Equal(t, "bca", Rotate("abc", -4))
	assert.Equal(t, "", Rotate("", -4))
}

func TestStartsWith(t *testing.T) {
	assert.False(t, StartsWith("", "", false))
	assert.False(t, StartsWith("", "abc", false))
	assert.False(t, StartsWith("", "ABC", true))
	assert.False(t, StartsWith("abcde", "", false))
	assert.False(t, StartsWith("ABC", "abc", false))
	assert.True(t, StartsWith("ABC", "abc", true))
	assert.True(t, StartsWith("a", "abc", true))
	assert.True(t, StartsWith("A", "abc", true))
}
