package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
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

func FindPythagoreanTriplets(v []int) []int {
	for i := 0; i < len(v); i++ {
		for k := 0; k < len(v); k++ {
			if k == i {
				continue
			}
			for l := 0; l < len(v); l++ {
				if l == i || l == k {
					continue
				}
				if math.Pow(float64(v[i]), 2)+math.Pow(float64(v[k]), 2) == math.Pow(float64(v[l]), 2) {
					return []int{v[i], v[k], v[l]}
				}
			}
		}
	}

	return []int{}
}

func Distance(w0 string, w1 string) int {
	// sitting
	// biting

	r0 := []rune(w0)
	r1 := []rune(w1)
	d := 0

	for i := 0; i < len(r0); i++ {
		if len(r0) < len(r1) {
			r0, r1 = r1, r0
		}

		w0 = string(r0)
		w1 = string(r1)

		var c0, c1, n0, n1 string

		n1 = n1

		if i < len(r0) {
			c0 = string(r0[i])
		}
		if i < len(r1) {
			c1 = string(r1[i])
		}

		if c0 == c1 {
			continue
		}

		if i+1 < len(r0) {
			n0 = string(r0[i+1])
		}
		if i+1 < len(r1) {
			n1 = string(r1[i+1])
		}

		if len(r0) > len(r1) {
			if n0 == c1 {
				r1 = append(append(r1[:0], r0[:1]...), r1[1:]...)

				d++
				continue
			}
		}

		d++
	}

	return d
}

func hash(s string) []byte {
	h := make([]byte, 26)

	a_ord := []byte("a")[0]
	for _, b := range []byte(s) {
		key := b - a_ord
		h[key] = h[key] + 1
	}

	return h
}

func Rnd(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}

	return int(nBig.Int64())
}

func FindSubstring(input string, search string) string {
	if len(search) > len(input) {
		return ""
	}

	searchHash := hash(search)

	for i := 0; i <= (len(input) - len(search)); i++ {
		substring := input[i : i+len(search)]
		currentHash := hash(substring)
		if bytes.Equal(searchHash, currentHash) {
			return substring
		}
	}

	return ""
}

func main() {
	fmt.Printf("run the test, not the main")
}
