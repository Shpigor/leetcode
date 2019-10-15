package leetcode

import (
	"testing"
)

const (
	anyCharacter      = '.'
	zeroOrMoreSymbols = '*'
)

const (
	INIT    = 0
	EMPTY   = 1
	ON_DOT  = 2
	ON_STAR = 3
	ON_CHAR = 4
)

type Pattern struct {
	pattern string
	size    int
	cursor  int
	state   int
}

func parse(pattern string) *Pattern {
	size := len(pattern)
	state := INIT
	if size == 0 {
		state = EMPTY
	}
	return &Pattern{
		pattern: pattern,
		size:    size,
		state:   state,
		cursor:  -1,
	}
}

func (p *Pattern) matchEmptyLine() bool {
	if p.state == EMPTY {
		return true
	} else {
		for hasNext, char := p.getNext(); hasNext; hasNext, char = p.getNext() {
			prevState := p.state
			switch char {
			case anyCharacter:
				p.state = ON_DOT
			case zeroOrMoreSymbols:
				p.state = ON_STAR
			default:
				p.state = ON_CHAR
			}
			if p.cursor == 0 && p.state == ON_DOT {
				return false
			} else if p.cursor == 1 {
				if p.state == ON_DOT || p.state == ON_CHAR {
					return false
				} else if p.state == ON_STAR && prevState == ON_CHAR {
					continue
				}
			} else if p.cursor > 1 {
				return false
			}
		}
	}
	return true
}

func (p *Pattern) getNext() (bool, rune) {
	if p.cursor < p.size-1 {
		p.cursor = p.cursor + 1
		return true, rune(p.pattern[p.cursor])
	}
	return false, -1
}

func isMatch(line string, pattern string) bool {
	ptr := parse(pattern)
	if len(line) == 0 && len(pattern) == 0 {
		return true
	} else if len(line) == 0 {
		return ptr.matchEmptyLine()
	}
	return false
}

func TestSimpleRegExp(t *testing.T) {
	line := "aa"
	pattern := "a"
	if isMatch(line, pattern) {
		t.Fail()
	}
}

func TestEmptyLineRegExp(t *testing.T) {
	line := ""
	pattern := "a*"
	if !isMatch(line, pattern) {
		t.Fail()
	}
}

func TestStarRegExp(t *testing.T) {
	line := "aa"
	pattern := "a*"
	if !isMatch(line, pattern) {
		t.Fail()
	}
}

func TestDotStarRegExp(t *testing.T) {
	line := "aa"
	pattern := ".*"
	if !isMatch(line, pattern) {
		t.Fail()
	}
}

func TestMultiStarRegExp(t *testing.T) {
	line := "aab"
	pattern := "c*a*b"
	if !isMatch(line, pattern) {
		t.Fail()
	}
}

func TestInvalidMultiStarRegExp(t *testing.T) {
	line := "mississippi"
	pattern := "mis*is*p*."
	if isMatch(line, pattern) {
		t.Fail()
	}
}
