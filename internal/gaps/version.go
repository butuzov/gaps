package gaps

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type VersionFlag bool

func (v VersionFlag) Decode(_ *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                       { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Fprintf(app.Stdout, "%s version: %s\n", app.Model.Name, vars["version"])
	app.Exit(0)
	return nil
}
