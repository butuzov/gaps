package main

import (
	"runtime/debug"

	"github.com/alecthomas/kong"

	"github.com/butuzov/gaps/internal/gaps"
)

func main() {
	var opts []kong.Option

	opts = append(opts, kong.Name("gaps"))
	opts = append(opts, kong.Description("mind te gaps"))
	opts = append(opts, kong.Vars{
		"version": func() string {
			version := "(devel)"
			if info, ok := debug.ReadBuildInfo(); ok {
				version = info.Main.Version
			}
			return version
		}(),
	})

	ctx := kong.Parse(&gaps.Settings{}, opts...)
	if err := ctx.Run(); err != nil {
		ctx.FatalIfErrorf(err)
	}
}
