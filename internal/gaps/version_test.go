package gaps

import (
	"bytes"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	v := VersionFlag(true)
	assert.True(t, v.IsBool())

	var bbOut, bbErr bytes.Buffer
	ka := &kong.Kong{
		Stdout: &bbOut,
		Stderr: &bbErr,
		Model: &kong.Application{
			Node: &kong.Node{Name: "gaps"},
		},
	}

	assert.Panics(t, func() {
		v.BeforeApply(ka, kong.Vars{"version": "v1.1.1"})
	})

	assert.Equal(t, "gaps version: v1.1.1\n", bbOut.String())
}
