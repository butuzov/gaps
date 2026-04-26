package gaps

import (
	"errors"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

// TODO: convert Args to a Settings.
// INFO(butuzov): RTFM https://github.com/alecthomas/kong#supported-tags.

type Settings struct {
	// Profile flag to specify the coverage file to analyze. By default,
	// coverage.out.
	Profile string  `type:"path" short:"c" default:"coverage.out" help:"coverage file"`
	Golden  *string `type:"path" short:"g"                        help:"coverage file to compare against"`
	Watch   bool    `            short:"w"                        help:"watch changes in coverrage profile, turns tui"`

	// --- Service flags -------------------------------------------------------
	Version VersionFlag `help:"Show version and exit." short:"v"`

	// Testing related settings.
	testing         bool
	testingExit     bool
	testingExitCode int
}

func (s *Settings) Testing(state bool) {
	s.testing = state
}

func (s *Settings) Exit(code int) {
	s.testingExitCode = code
	s.testingExit = true
}

func (s *Settings) Run(ctx *kong.Context) error {
	_ = ctx

	if s.testing && s.testingExit {
		if s.testingExitCode > 0 {
			return fmt.Errorf("exit code: %d", s.testingExitCode)
		}
		return nil
	}

	return New(s).Run()
}

// TODO: func (s *Settings) Help() string

func (s *Settings) Validate() error {
	if _, err := os.Stat(s.Profile); errors.Is(err, os.ErrNotExist) {
		return ErrProfileNotFound.WithMeta(fmt.Sprintf("argument --profile=%s", s.Profile))
	}

	if s.Golden != nil {
		if _, err := os.Stat(*s.Golden); errors.Is(err, os.ErrNotExist) {
			return ErrGoldenNotFound.WithMeta(fmt.Sprintf("argument --golden=%s", *s.Golden))
		}
	}

	return nil
}

func (s *Settings) mode() Mode {
	var m Mode

	if s.Watch {
		m |= ModeWatch
	}

	// if no golden, only analyze enabled.
	if s.Golden != nil {
		m |= ModeCompare
	} else {
		m |= ModeAnalize
	}

	// Watch not enabled, turn on fire and forgot
	if !m.Has(ModeWatch) {
		m |= ModeFireAndForgot
	}

	return m
}
