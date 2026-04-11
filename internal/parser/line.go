package parser

import (
	"errors"
	"fmt"
)

type Line struct {
	File  string // File name for this block.
	Block Block  //
	Count uint32 // Number of times this block was executed.
}

type Block struct {
	Line0 uint32 // Line number for block start.
	Col0  uint32 // Column number for block start.
	Line1 uint32 // Line number for block end.
	Col1  uint32 // Column number for block end.
	Stmts uint32 // Number of statements included in this block.
}

// parseLine is fast enought go coverage line parser, better version then
// stateMachineParserParser (x23 times)
func parseLine(line string) (l Line, err error) {
	if len(line) == 0 {
		return l, errors.New("empty line")
	}

	// Searching for file....
	colonIdx := -1
	for i := 0; i < len(line); i++ {
		if line[i] == ':' {
			colonIdx = i
			break
		}
	}

	if colonIdx < 0 || colonIdx >= len(line) {
		return l, errors.New("out of bounds check fail")
	}

	l.File = line[:colonIdx]
	// Checking blocks (numbers)

	length := len(line)
	parseUntil := func(start int, delim byte) (count uint32, index int, err error) {
		if start < 0 || start >= length {
			return 0, 0, errors.New("delimiter not found")
		}

		num := uint32(0)
		found := false
		for i := start; i < length; i++ {
			c := line[i]

			if c == delim {
				return num, i + 1, nil
			}

			if c >= '0' && c <= '9' {
				num = num*10 + uint32(c-'0')
				found = true
				continue
			}
			return 0, 0, errors.New("invalid digit")
		}
		if !found {
			return 0, 0, errors.New("delimiter not found")
		}

		return 0, 0, errors.New("delimiter not found")
	}

	index := colonIdx + 1

	l.Block.Line0, index, err = parseUntil(index, '.')
	if err != nil {
		return l, fmt.Errorf("parsing line0: %v", err)
	}

	l.Block.Col0, index, err = parseUntil(index, ',')
	if err != nil {
		return l, fmt.Errorf("parsing col0: %v", err)
	}

	l.Block.Line1, index, err = parseUntil(index, '.')
	if err != nil {
		return l, fmt.Errorf("parsing line1: %v", err)
	}

	l.Block.Col1, index, err = parseUntil(index, ' ')
	if err != nil {
		return l, fmt.Errorf("parsing col1: %v", err)
	}

	l.Block.Stmts, index, err = parseUntil(index, ' ')
	if err != nil {
		return l, fmt.Errorf("parsing stmts: %v", err)
	}

	if index < 0 || index >= length {
		return l, err
	}

	for i := index; i < length; i++ {
		c := line[i]
		if c >= '0' && c <= '9' {
			l.Count = l.Count*10 + uint32(c-'0')
			continue
		}
		return l, errors.New("parsing count: unexpected stop")
	}

	return l, nil
}
