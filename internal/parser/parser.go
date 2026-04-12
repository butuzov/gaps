package parser

import (
	"bufio"
	"errors"
	"io"
	"path/filepath"
	"strings"
)

// Profile is
type Profile struct {
	Mode     string              // atomic/set/count
	Counters map[string][]uint32 // counters
	Blocks   map[string][]Block  // covered blocks
	Packages map[string][]string // packages with files
}

func parser(r io.Reader) (Profile, error) {
	scanner := bufio.NewScanner(r)

	profile := Profile{
		Counters: map[string][]uint32{},
		Blocks:   map[string][]Block{},
		Packages: map[string][]string{},
	}

	if scanner.Scan() {
		mode, ok := strings.CutPrefix(scanner.Text(), "mode: ")
		if !ok {
			return profile, errors.New("mode is missing")
		}
		profile.Mode = mode
	}

	for scanner.Scan() {
		res, err := parseLine(scanner.Text())
		if err != nil {
			continue
		}

		profile.Counters[res.File] = append(profile.Counters[res.File], res.Count)
		profile.Blocks[res.File] = append(profile.Blocks[res.File], res.Block)

		pkg, file := filepath.Split(res.File)
		profile.Packages[pkg] = append(profile.Packages[pkg], file)
	}

	return profile, nil
}
