package gaps

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (gaps *Gaps) Init() tea.Cmd {
	return nil
}

func (gaps *Gaps) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, gaps.keysHelp.Quit):
			return gaps, tea.Quit
		case key.Matches(msg, gaps.keysHelp.Help):
			gaps.help.ShowAll = !gaps.help.ShowAll
		}
	case tea.WindowSizeMsg:
		gaps.width = msg.Width
		gaps.height = msg.Height
	}
	return gaps, nil
}

func (gaps *Gaps) View() tea.View {
	header, footer := gaps.headerView(), gaps.footerView()

	content := lipgloss.NewStyle().
		Width(gaps.width).
		Height(gaps.height-lipgloss.Height(header)-lipgloss.Height(footer)).
		Align(lipgloss.Left, lipgloss.Center).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Render("content")

	v := tea.NewView(lipgloss.JoinVertical(
		lipgloss.Left,
		// TODO: correct template
		header,
		content,
		footer,
	))

	v.AltScreen = true
	return v
}

func (gaps *Gaps) headerView() string {
	return lipgloss.NewStyle().
		Width(gaps.width).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Render("header")
}

func (gaps *Gaps) footerView() string {
	return gaps.help.View(gaps.keysHelp)
}
