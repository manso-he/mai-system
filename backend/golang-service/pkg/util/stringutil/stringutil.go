package stringutil

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"manso.live/backend/golang-service/pkg/constant"
)

const Empty = ""

func Equals(s1, s2 string) bool {
	return s1 == s2
}

func EqualsIgnoreCase(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

func NotEquals(s1, s2 string) bool {
	return !Equals(s1, s2)
}

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == Empty
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func IsEmpty(s string) bool {
	return s == Empty
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func Length(s string) int {
	return utf8.RuneCountInString(s)
}

func Contains(str, searchStr string) bool {
	return strings.Contains(str, searchStr)
}

func ContainsAnyString(src string, candidates ...string) bool {
	for _, str := range candidates {
		if Contains(src, str) {
			return true
		}
	}
	return false
}

func StartsWith(str, prefix string) bool {
	return startsWithIgnoreCase(str, prefix, false)
}

func StartsWithIgnoreCase(str, prefix string) bool {
	return startsWithIgnoreCase(str, prefix, true)
}

func startsWithIgnoreCase(str, prefix string, ignore bool) bool {
	if len(prefix) > len(str) {
		return false
	}
	s := str[:len(prefix)]
	if ignore {
		return strings.EqualFold(s, prefix)
	}
	return s == prefix
}

func Trim(s, cutset string) string {
	return strings.Trim(s, cutset)
}

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// Concat combines multiple strings into a single string
func Concat(str ...string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}

// SlashJoin concatenates the elements of a to create a single string, separated by slashes.
//
// e.g., SlashJoin("a", "b", "c/d") => "a/b/c/d"
func SlashJoin(str ...string) string {
	return strings.Join(str, "/")
}

func IfBlank(str1, str2 string) string {
	if !IsBlank(str1) {
		return str1
	}
	return str2
}

func SingleSlashJoin(str ...string) string {
	trimmedStr := make([]string, 0)
	for _, s := range str {
		trimmed := strings.Trim(s, "/")
		if IsEmpty(trimmed) {
			continue
		}
		trimmedStr = append(trimmedStr, trimmed)
	}
	return strings.Join(trimmedStr, "/")
}

func PeriodJoin(str ...string) string {
	return strings.Join(str, ".")
}

// HasUppercase checks if the specified string contains uppercase characters
func HasUppercase(str string) (has bool) {
	for _, s := range str {
		if unicode.IsUpper(s) {
			return true
		}
	}
	return false
}

// HasLowercase checks if the specified string contains lowercase characters
func HasLowercase(str string) (has bool) {
	for _, s := range str {
		if unicode.IsLower(s) {
			return true
		}
	}
	return false
}

func SplitCommaAndToIntSlice(str string) ([]int, error) {
	if str == "" {
		return []int{}, nil
	}
	strs := strings.Split(str, constant.Comma)
	ints := make([]int, len(strs))
	for i, v := range strs {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ints[i] = n
	}
	return ints, nil
}
