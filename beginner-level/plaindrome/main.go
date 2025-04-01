package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []rune
}

func (s *Stack) Pop(item rune) (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	top := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return top, true

}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func main() {

	newStr := "madam"
	

	fmt.Println(" Is this Plaindrome:::::", isPlaindrome(newStr))

}

func isPlaindrome(abc string) bool {
	stk := Stack{}
	
	abc = strings.ToLower(strings.ReplaceAll(abc, " ", ""))

	for _, char:= range abc {
		stk.Push(char)
	}

	for _,a := range abc {
		top,_ := stk.Pop(a)

		if a != top{
			return false
		}
			
	}

	return true

}