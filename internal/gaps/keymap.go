package gaps

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	Filter key.Binding
	Help   key.Binding
	Quit   key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help}, // One Per Column
		{k.Quit},
		{k.Filter},
	}
}

var keys = keyMap{
	Filter: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "filter"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}

func Help() help.Model {
	h := help.New()
	h.FullSeparator = " • "
	return h
}
