package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestLengthOfLongestSubstring5(t *testing.T) {
	require.Equal(t, 10, lengthOfLongestSubstring("abrkaabcdefghijjxxx"))
}

func TestLongestPalindrome(t *testing.T) {
	require.Equal(t, "a", longestPalindrome("a"))
	require.Equal(t, "aa", longestPalindrome("aa"))
	require.Equal(t, "anana", longestPalindrome("banana"))
	require.Equal(t, "illi", longestPalindrome("million"))
}

func TestIsValidParentheses(t *testing.T) {
	require.True(t, isValidParentheses(""))
	require.False(t, isValidParentheses("}"))
	require.False(t, isValidParentheses("(}"))
	require.True(t, isValidParentheses("()"))
	require.True(t, isValidParentheses("((()))"))
	require.True(t, isValidParentheses("[()]{}"))
	require.False(t, isValidParentheses("({[)]"))
}

func TestBinarySearch(t *testing.T) {
	require.Equal(t, -1, BinarySearch([]int{}, 0))
	require.Equal(t, 0, BinarySearch([]int{0}, 0))
	require.Equal(t, -1, BinarySearch([]int{1}, 0))
	require.Equal(t, 0, BinarySearch([]int{0, 1}, 0))
	require.Equal(t, 1, BinarySearch([]int{0, 1}, 1))
	require.Equal(t, 2, BinarySearch([]int{0, 1, 2}, 2))
	require.Equal(t, -1, BinarySearch([]int{0, 1, 2}, 3))
	require.Equal(t, -1, BinarySearch([]int{0, 1, 3}, 2))
}

func TestRange(t *testing.T) {
	require.EqualValues(t, IntsToString([]int{6, 8}), IntsToString(Range([]int{1, 3, 3, 5, 7, 8, 9, 9, 9, 15}, 9)))
	require.EqualValues(t, IntsToString([]int{1, 2}), IntsToString(Range([]int{100, 150, 150, 153}, 150)))
	require.EqualValues(t, IntsToString([]int{-1, -1}), IntsToString(Range([]int{1, 2, 3, 4, 5, 6, 10}, 9)))
}

func TestFindPythagoreanTriplets(t *testing.T) {
	require.EqualValues(t, []int{12, 5, 13}, FindPythagoreanTriplets([]int{3, 12, 5, 13}))
}

/*
func TestDistance(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		//{"otto", "otto", 0},
		//{"abc", "abx", 1},
		//{"abc", "xbc", 1},
		//{"sitting", "biting", 2},
		//{"abc", "123", 3},
		//{"a", "", 1},
		//{"ab", "", 2},
		//{"", "hello", 5},
		//{"hello", "", 5},
		//{"hello", "hello", 0},
		//{"ab", "aa", 1},
		//{"ab", "ba", 2},
		//{"ab", "aaa", 2},
		//{"bbb", "a", 3},
		//{"kitten", "sitting", 3},
		{"distance", "difference", 5},
		{"levenshtein", "frankenstein", 6},
		{"resume and cafe", "resumes and cafes", 2},
	}
	for _, d := range tests {
		fmt.Printf("%+v\n", d)
		require.Equal(t, d.want, Distance(d.a, d.b))
	}
}
*/

func RndString(len int) string {
	s := strings.Builder{}
	for i := 0; i < len; i++ {
		s.WriteByte(97 + byte(Rnd(26)))
	}

	return s.String()
}

func TestFindSubstring(t *testing.T) {
	require.Equal(t, "zyx", FindSubstring("xyyzyzyx", "xyz"))

	search := RndString(30)
	input := RndString(470) + search

	require.Equal(t, search, FindSubstring(input, search))
}

func BenchmarkFindSubstring(b *testing.B) {
	search := RndString(30)
	input := RndString(470) + search

	FindSubstring(input, search)
}
