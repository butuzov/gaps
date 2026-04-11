package parser

// info(butuzov): this file contains other implementations of line parsing, slow but clear
//                more complex but faster, both used to verify corectness and mesure speed.

import (
	"errors"
	"regexp"
	"strconv"
)

var profileLineRe = regexp.MustCompile(`^(.+):([0-9]+)\.([0-9]+),([0-9]+)\.([0-9]+) ([0-9]+) ([0-9]+)$`)

// regexpTestParserTest Implementation (1) - baseline implementation.
func regexpTestParser(line string) (l Line, err error) {
	matches := profileLineRe.FindStringSubmatch(line)
	if matches == nil {
		return l, errors.New("no matches")
	}

	return Line{
		File: matches[1],
		Block: Block{
			Line0: uint32(toInt(matches[2])),
			Col0:  uint32(toInt(matches[3])),
			Line1: uint32(toInt(matches[4])),
			Col1:  uint32(toInt(matches[5])),
			Stmts: uint32(toInt(matches[6])),
		},
		Count: uint32(toInt(matches[7])),
	}, nil
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// stateMachineParserParser Test Implementation (2) - one more baseline implementation.
func stateMachineParserParser(line string) (l Line, err error) {
	if len(line) == 0 {
		return l, errors.New("empty line")
	}

	var state int
	for i := 0; i < len(line); i++ {
		c := line[i]

		switch state {
		case 0: // file

			if c == ':' {
				state = 1
				l.File = line[:i]
			}

		case 1: // Line0

			if c >= '0' && c <= '9' {
				l.Block.Line0 = l.Block.Line0*10 + uint32(c-'0')
				continue
			}

			if c == '.' {
				state = 2
				continue
			}
			return l, errors.New("unexpeced stop at parsing line0")

		case 2: // Col0

			if c >= '0' && c <= '9' {
				l.Block.Col0 = l.Block.Col0*10 + uint32(c-'0')
				continue
			}

			if c == ',' {
				state = 3
				continue
			}
			return l, errors.New("unexpeced stop at parsing col0")

		case 3: // Line1

			if c >= '0' && c <= '9' {
				l.Block.Line1 = l.Block.Line1*10 + uint32(c-'0')
				continue
			}

			if c == '.' {
				state = 4
				continue
			}
			return l, errors.New("unexpeced stop at parsing line1")

		case 4: // Col1

			if c >= '0' && c <= '9' {
				l.Block.Col1 = l.Block.Col1*10 + uint32(c-'0')
				continue
			}

			if c == ' ' {
				state = 5
				continue
			}
			return l, errors.New("unexpeced stop at parsing col1")

		case 5: // Stamts

			if c >= '0' && c <= '9' {
				l.Block.Stmts = l.Block.Stmts*10 + uint32(c-'0')
				continue
			}

			if c == ' ' {
				state = 6
				continue
			}
			return l, errors.New("unexpeced stop at parsing statments")

		case 6: // Count

			if c >= '0' && c <= '9' {
				l.Count = l.Count*10 + uint32(c-'0')
				continue
			}

			return l, errors.New("unexpeced stop at parsing count")
		}
	}

	return l, nil
}
