package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/require"

	"github.com/butuzov/gaps/internal/gaps"
)

var update = flag.Bool("update", false, "update golden files")

func TestMain(m *testing.M) {
	flag.Parse()

	if *update {
		os.RemoveAll(filepath.Join("testdata", "fixtures"))
		err := os.Mkdir(filepath.Join("testdata", "fixtures"), 0o774)
		if err != nil {
			panic(fmt.Sprintf("failed to create fixtures dir: %s", err))
		}
	}

	os.Exit(m.Run())
}

// golden returns a path to golden file, so we can avoid passing path var
func golden(t *testing.T, name string, got []byte) []byte {
	t.Helper()

	path := filepath.Join("testdata", "fixtures", name)

	if *update {
		require.NoError(t, os.WriteFile(path, got, 0o664))
	}

	want, err := os.ReadFile(path)
	require.NoError(t, err)

	return want
}

func runTestGaps(bb *bytes.Buffer, args []string, opts ...kong.Option) error {
	var options []kong.Option
	settings := &gaps.Settings{}
	settings.Testing(true)

	options = append(options, kong.Writers(bb, bb))
	options = append(options, kong.Name("gaps"))
	options = append(options, kong.Description("mind te gaps"))
	options = append(options, kong.Vars{
		"version": func() string {
			version := "(devel)"
			if info, ok := debug.ReadBuildInfo(); ok {
				version = info.Main.Version
			}
			return version
		}(),
	})
	options = append(options, kong.Exit(func(code int) {
		settings.Exit(code)
	}))
	options = append(options, opts...)

	app, err := kong.New(settings, options...)
	if err != nil {
		return err
	}

	ctx, err := app.Parse(args)
	if err != nil {
		return err
	}

	return ctx.Run()
}
