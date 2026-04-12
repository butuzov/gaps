package parser

import (
	"testing"
)

var (
	sink Line
	err  error
)

func BenchmarkCovLineParse(b *testing.B) {
	tesbenhLines := []string{
		"github.com/butuzov/ireturn/analyzer/analyzer.go:273.27,275.3 12 11",
		"github.com/golangci/golangci-lint/v2/internal/x/tools/diff/lcs/old.go:311.37,315.21 3 115030",
		"github.com/ollama/ollama/anthropic/anthropic.go:321.2,321.33 1 15",
		"github.com/junegunn/fzf/src/util/atexit.go:12.15,13.39 1 0",
	}

	callbacks := []struct {
		name string
		call func(string) (Line, error)
	}{
		{"test_re", regexpTestParser},
		{"test_sm", stateMachineParserParser},
		{"parser", parseLine},
	}

	for _, impl := range callbacks {
		b.Run(impl.name, func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				sink, err = impl.call(tesbenhLines[n%len(tesbenhLines)])
			}
		})
	}
}
