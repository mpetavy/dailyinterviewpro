package main

import (
	"fmt"
)

//Hi, here's your problem today. This problem was recently asked by Microsoft:
//
//Given a string, find the length of the longest substring without repeating characters.
//
//Here is an example solution in Python language. (Any language is OK to use in an interview, though we'd recommend Python as a generalist language utilized by companies like Google, Facebook, Netflix, Dropbox, Pinterest, Uber, etc.,)

func lengthOfLongestSubstring(s string) int {
	var lastR rune
	max := 0
	c := 0

	for i, r := range s {
		c++

		if max < c {
			max = c
		}

		if i == 0 {
			continue
		}

		if lastR == r {
			c = 0
		}

		lastR = r
	}

	return max
}

//Hi, here's your problem today. This problem was recently asked by Twitter:
//
//A palindrome is a sequence of characters that reads the same backwards and forwards. Given a string, s, find the longest palindromic substring in s.
//
//Example:
//Input: "banana"
//Output: "anana"
//
//Input: "million"
//Output: "illi"

func longestPalindrome(s string) string {
	r := ""

	for i := 1; i <= len(s); i++ {
		for k := 0; k < i; k++ {
			sub := string(s[k:i])
			rsub := Reverse(sub)
			if sub == rsub {
				if len(r) < len(sub) {
					r = sub
				}
			}
		}
	}

	return r
}

func indexAny(s string, v string) int {
	for _, ch := range s {
		for i, f := range v {
			if ch == f {
				return i
			}
		}
	}

	return -1
}

//Hi, here's your problem today. This problem was recently asked by Uber:
//
//Imagine you are building a compiler. Before running any code, the compiler must check that the parentheses in the program are balanced. Every opening bracket must have a corresponding closing bracket. We can approximate this using strings.
//
//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//An input string is valid if:
//- Open brackets are closed by the same type of brackets.
//- Open brackets are closed in the correct order.
//- Note that an empty string is also considered valid.
//
//Example:
//Input: "((()))"
//Output: True
//
//Input: "[()]{}"
//Output: True
//
//Input: "({[)]"
//Output: False

func isValidParentheses(s string) bool {
	const (
		openBrackets  = "({["
		closeBrackets = ")}]"
	)

	p := make([]int, 0)
	for _, ch := range s {
		index := indexAny(string(ch), openBrackets)
		if index != -1 {
			p = append(p, index)

			continue
		}

		index = indexAny(string(ch), closeBrackets)
		if index != -1 {
			if len(p) == 0 {
				return false
			}

			if index != p[len(p)-1] {
				return false
			}

			p = p[0 : len(p)-1]
		}
	}

	return len(p) == 0
}

//Hi, here's your problem today. This problem was recently asked by AirBNB:
//
//Given a sorted array, A, with possibly duplicated elements, find the indices of the first and last occurrences of a target element, x. Return -1 if the target is not found.
//
//Example:
//Input: A = [1,3,3,5,7,8,9,9,9,15], target = 9
//Output: [6,8]
//
//Input: A = [100, 150, 150, 153], target = 150
//Output: [1,2]
//
//Input: A = [1,2,3,4,5,6,10], target = 9
//Output: [-1, -1]

func Range(arr []int, b int) []int {
	index := BinarySearch(arr, b)

	left := -1
	right := -1

	if index == -1 {
		return []int{left, right}
	}

	for i := index; i >= 0 && b <= arr[i]; i-- {
		if b == arr[i] {
			left = i
		}
	}

	for i := index; i < len(arr) && b >= arr[i]; i++ {
		if b == arr[i] {
			right = i
		}
	}

	return []int{left, right}
}

func main() {
	fmt.Printf("run the test, not the main")
}
