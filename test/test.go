package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome4("aacabdkacaa"))
}

func longestPalindrome4(s string) string {
	n, m, i, j := 0, 0, 0, 0
	for k := 0; ; k++ {
		if j+1 == len(s) {
			break
		}
		if k%2 == 0 {
			i = k/2 + 1
			j = k / 2
		} else {
			i, j = k/2+1, k/2+1
		}
		if s[i] != s[j] {
			continue
		}
		for {
			if i+1 == len(s) || j == 0 {
				break
			}
			if s[i+1] == s[j-1] {
				i++
				j--
				continue
			} else {
				break
			}

		}
		if i-j > m-n {
			m = i
			n = j
		}
	}
	return s[n : m+1]
}
