package gaps

import (
	"bytes"
	"fmt"

	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"

	"github.com/butuzov/gaps/internal/parser"
)

type Gaps struct {
	Settings Settings

	// Application
	chProfile chan parser.Profile
	chGolden  chan parser.Profile

	// TUI ---------------------------------------------------------------------
	width, height int  // Terminal Dimensions
	mode          Mode // Application Mode
	help          help.Model
	keysHelp      keyMap

	tmp string
}

func New(settings *Settings) *Gaps {
	gaps := &Gaps{
		Settings: *settings,
		help:     Help(),
		keysHelp: keys,
	}

	gaps.mode = settings.mode()
	gaps.chProfile = make(chan parser.Profile)
	if gaps.mode.Has(ModeCompare) {
		gaps.chGolden = make(chan parser.Profile)
	}

	// TMP
	var bb bytes.Buffer
	// spew.Fdump(&bb, gaps)
	fmt.Fprintf(&bb, "coverage(cur): %s\n", settings.Profile)
	if settings.Golden != nil {
		fmt.Fprintf(&bb, "coverage(gld): %s\n", *settings.Golden)
	}

	gaps.tmp = bb.String()

	return gaps
}

func (gaps *Gaps) Run() error {
	p := tea.NewProgram(gaps)
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
