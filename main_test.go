package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring5(t *testing.T) {
	assert.Equal(t, 10, lengthOfLongestSubstring("abrkaabcdefghijjxxx"))
}

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "a", longestPalindrome("a"))
	assert.Equal(t, "aa", longestPalindrome("aa"))
	assert.Equal(t, "anana", longestPalindrome("banana"))
	assert.Equal(t, "illi", longestPalindrome("million"))
}

func TestIsValidParentheses(t *testing.T) {
	assert.True(t, isValidParentheses(""))
	assert.False(t, isValidParentheses("}"))
	assert.False(t, isValidParentheses("(}"))
	assert.True(t, isValidParentheses("()"))
	assert.True(t, isValidParentheses("((()))"))
	assert.True(t, isValidParentheses("[()]{}"))
	assert.False(t, isValidParentheses("({[)]"))
}

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, -1, BinarySearch([]int{}, 0))
	assert.Equal(t, 0, BinarySearch([]int{0}, 0))
	assert.Equal(t, -1, BinarySearch([]int{1}, 0))
	assert.Equal(t, 0, BinarySearch([]int{0, 1}, 0))
	assert.Equal(t, 1, BinarySearch([]int{0, 1}, 1))
	assert.Equal(t, 2, BinarySearch([]int{0, 1, 2}, 2))
	assert.Equal(t, -1, BinarySearch([]int{0, 1, 2}, 3))
	assert.Equal(t, -1, BinarySearch([]int{0, 1, 3}, 2))
}

func TestRange(t *testing.T) {
	assert.EqualValues(t, IntsToString([]int{6, 8}), IntsToString(Range([]int{1, 3, 3, 5, 7, 8, 9, 9, 9, 15}, 9)))
	assert.EqualValues(t, IntsToString([]int{1, 2}), IntsToString(Range([]int{100, 150, 150, 153}, 150)))
	assert.EqualValues(t, IntsToString([]int{-1, -1}), IntsToString(Range([]int{1, 2, 3, 4, 5, 6, 10}, 9)))
}

func TestFindPythagoreanTriplets(t *testing.T) {
	assert.EqualValues(t, []int{12, 5, 13}, FindPythagoreanTriplets([]int{3, 12, 5, 13}))
}

func TestDistance(t *testing.T) {
	assert.EqualValues(t, 0, Distance("otto", "otto"))
	assert.EqualValues(t, 1, Distance("abc", "abx"))
	assert.EqualValues(t, 1, Distance("abc", "xbc"))
	assert.EqualValues(t, 2, Distance("sitting", "biting"))
	assert.EqualValues(t, 3, Distance("abc", "123"))
}
