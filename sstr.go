package sstr

import (
	"strings"
)

type Sstr struct {
	Str string
}

// 返回值函数
func (s *Sstr) Val() string { return s.Str }

// 去除所有空格
func (s *Sstr) RemoveAllSpaces() {
	strings.TrimSpace(s.Str)
	s.Str = strings.Join(strings.Split(s.Str, " "), "")
}

// 清除字符串
func (s *Sstr) Clean() {
	s.Str = ""
}

// 保留字符串中的数字
func (s *Sstr) KeepNumbers() {
	s.Str = strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s.Str)
}

// 保留字符串中的大写字符
func (s *Sstr) KeepUpperChar() {
	s.Str = strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z'{
			return r
		}
		return -1
	}, s.Str)
}

// 保留字符串中的小写字符
func (s *Sstr) KeepLowerChar() {
	s.Str = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z'{
			return r
		}
		return -1
	}, s.Str)
}

// 保留特定字符 可变参数 第一个保留数字 第二个保留大写字母 第三个保留小写字母 可以随意组合)
func (s *Sstr) KeepSpecificChar(args ... bool) {
	var Number, Upper, Lower bool
	if len(args) == 1 {
		Number = args[0]
	}else if len(args) == 2 {
		Number = args[0]
		Upper = args[1]
	}else if len(args) == 3 {
		Number = args[0]
		Upper = args[1]
		Lower = args[2]
	}
	s.Str = strings.Map(func(r rune) rune {
		if Number{
			if r >= '0' && r <= '9' {
				return r
			}
		}
		if Upper {
			if r >= 'A' && r <= 'Z'{
				return r
			}
		}
		if Lower {
			if r >= 'a' && r <= 'z'{
				return r
			}
		}
		return -1
	}, s.Str)
}

