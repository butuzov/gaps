package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	var bb bytes.Buffer
	err := runTestGaps(&bb, []string{"-v"})

	want := golden(t, "version.txt", bb.Bytes())
	require.NoError(t, err)
	assert.Equal(t, string(want), bb.String())
}

func TestHelp(t *testing.T) {
	var bb bytes.Buffer
	err := runTestGaps(&bb, []string{"--help"})

	want := golden(t, "help.txt", bb.Bytes())
	require.NoError(t, err)
	assert.Equal(t, string(want), bb.String())
}
