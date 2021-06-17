package main

import "fmt"

func lengthOfLongestSubstring1(s string, c chan string) {
	subStr := []int{0, 1}
	for key, _ := range s {
		hash, start, end := make(map[string]bool), key, key
		for end < len(s) && hash[string(s[end])] != true {
			hash[string(s[end])] = true
			end += 1
		}
		if subStr[1]-subStr[0] < end-start {
			subStr[0], subStr[1] = start, end
		}
	}
	if len(s) > 0 {
		c <- s[subStr[0]:subStr[1]]
	}
	c <- s[subStr[0]:subStr[1]]
}

func lengthOfLongestSubstring2(s string, c chan string) {
	seen := make(map[string]bool)
	left, right := 0, 0
	longest := []int{0, 0}
	for right < len(s) {
		if _, ok := seen[string(s[right])]; !ok {
			seen[string(s[right])] = true
			right++
		} else {
			delete(seen, string(s[left]))
			left++
		}
		if longest[1]-longest[0] < right-left {
			longest[1], longest[0] = right, left
		}
	}
	fmt.Println("2")
	c <- s[longest[0]:longest[1]]
}

func lengthOfLongestSubstring3(s string, c chan string) {
	usedChars := make([]rune, 0)
	var ans string = ""
	for _, char := range s {
		index := indexOf(usedChars, char)
		if index != -1 {
			if len(usedChars) > len(ans) {
				ans = string(usedChars)
			}
			usedChars = usedChars[index+1:]
		}
		usedChars = append(usedChars, char)
	}
	if len(ans) > len(usedChars) {
		c <- ans
	} else {
		c <- string(usedChars)
	}
}

func indexOf(chars []rune, char rune) int {
	for i, currChar := range chars {
		if char == currChar {
			return i
		}
	}
	return -1
}

func main() {
	c := make(chan string)
	go lengthOfLongestSubstring1("pwekwekkewi", c)
	go lengthOfLongestSubstring2("pwekwekkewi", c)
	go lengthOfLongestSubstring3("pwekwekkewi", c)
	for count := 0; count < 3; count++ {
		v := <-c
		fmt.Println(v)
	}
}
