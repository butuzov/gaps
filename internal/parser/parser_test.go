package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	t.Parallel()
	f, err := os.Open("testdata/golangci_lint_1001.out")
	require.NoError(t, err)
	defer f.Close() //nolint

	cp, err := parser(f)

	require.NoError(t, err)
	assert.Equal(t, "atomic", cp.Mode)

	// 39 records (files)
	assert.Len(t, cp.Counters, 39)
	assert.Len(t, cp.Blocks, 39)

	// Asserting Block Info
	{
		have := cp.Blocks["github.com/golangci/golangci-lint/v2/internal/go/mmap/mmap.go"]
		want := []Block{
			{25, 44, 27, 16, 2},
			{27, 16, 29, 3, 1},
			{30, 2, 31, 24, 2},
		}
		assert.Equal(t, want, have)
	}

	// asserting Count
	{
		have := cp.Counters["github.com/golangci/golangci-lint/v2/cmd/golangci-lint/main.go"]
		want := make([]uint32, 14)
		assert.Equal(t, want, have)
	}
}
