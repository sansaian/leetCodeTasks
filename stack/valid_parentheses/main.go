package main

import "fmt"

func main() {
	fmt.Println(isValid("]"))    //false
	fmt.Println(isValid("()"))   //true
	fmt.Println(isValid("([)]")) // false
	fmt.Println(isValid("{[]}")) // true
	fmt.Println(isValid("(])"))  //false
}

func isValid(s string) bool {
	st := stack{
		store: []rune{},
	}
	for _, ch := range s {
		switch {
		case ch == '(' || ch == '{' || ch == '[':
			st.store = append(st.store, ch)
		default:
			if len(st.store) == 0 {
				return false
			}
			if ok := st.remove(ch); !ok {
				return false
			}
		}
	}
	return len(st.store) == 0
}

type stack struct {
	store []rune
}

func (s *stack) remove(ch rune) bool {

	switch ch {
	case ')':
		if s.store[len(s.store)-1] != '(' {
			return false
		}

	case '}':
		if s.store[len(s.store)-1] != '{' {
			return false
		}

	case ']':
		if s.store[len(s.store)-1] != '[' {
			return false
		}
	}
	s.store = s.store[:len(s.store)-1]
	return true
}
