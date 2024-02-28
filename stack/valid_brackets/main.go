package main

import "fmt"

func main() {
	fmt.Println(check("()", [][]byte{{'(', ')'}}))
}

type stack struct {
	store []byte
}

type BracketConfig struct {
	openBrackets  map[byte]*struct{}
	closeBrackets map[byte]byte
}

func check(str string, config [][]byte) bool {
	parsedCfg := parseCfg(config)
	st := stack{
		store: []byte{},
	}

	for i := 0; i < len(str); i++ {
		switch {
		case parsedCfg.openBrackets[str[i]] != nil:
			st.store = append(st.store, str[i])
		default:
			if len(st.store) == 0 {
				return false
			}
			if ok := st.remove(str[i], parsedCfg); !ok {
				return false
			}
		}
	}
	return len(st.store) == 0
}

func (s *stack) remove(ch byte, parsedCfg *BracketConfig) bool {
	switch {
	case parsedCfg.closeBrackets[ch] != 0x0:
		if s.store[len(s.store)-1] != parsedCfg.closeBrackets[ch] {
			return false
		}
	}
	s.store = s.store[:len(s.store)-1]
	return true
}

func parseCfg(cfg [][]byte) *BracketConfig {
	brConfig := BracketConfig{
		openBrackets:  make(map[byte]*struct{}),
		closeBrackets: make(map[byte]byte),
	}
	for _, v := range cfg {
		brConfig.openBrackets[v[0]] = &struct{}{}
		brConfig.closeBrackets[v[1]] = v[0]
	}
	return &brConfig
}
