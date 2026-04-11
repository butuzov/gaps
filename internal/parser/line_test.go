package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLines(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, line string
		want       Line
	}{
		{
			"simple", "github.com/butuzov/gaps/main.go:3.13,5.2 1 0",
			Line{"github.com/butuzov/gaps/main.go", Block{3, 13, 5, 2, 1}, 0},
		},
		{
			"bigger_numbers", "github.com/butuzov/gaps/main.go:301.130,95.12 101 12",
			Line{"github.com/butuzov/gaps/main.go", Block{301, 130, 95, 12, 101}, 12},
		},
	}

	callbacks := []struct {
		name string
		call func(string) (Line, error)
	}{
		{"test_re", regexpTestParser},
		{"test_sm", stateMachineParserParser},
		{"current", parseLine},
	}

	for _, test := range tests {
		for _, testFunc := range callbacks {
			t.Run(testFunc.name+"/"+test.name, func(t *testing.T) {
				t.Parallel()

				have, err := testFunc.call(test.line)

				require.NoError(t, err)
				assert.Equal(t, test.want, have)
			})
		}
	}
}
