package home

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var selectKeyEnter = key.NewBinding(
	key.WithKeys(
		"enter",
	),
	key.WithHelp(
		"enter",
		"select",
	),
)

var selectKeyQuit = key.NewBinding(
	key.WithKeys(
		"q",
		"ctrl+c",
	),
	key.WithHelp(
		"q",
		"quit",
	),
)

var selectKeyUp = key.NewBinding(
	key.WithKeys(
		"up",
		"k",
	),
	key.WithHelp(
		"↑",
		"up",
	),
)

var selectKeyDown = key.NewBinding(
	key.WithKeys(
		"down",
		"j",
	),
	key.WithHelp(
		"↓",
		"down",
	),
)

type Model struct {
	choices []Option
	cursor  int
}

func (m *Model) currChoice() Option {
	return m.choices[m.cursor]
}

func InitialModel() Model {
	return Model{choices: options}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, selectKeyEnter):
			if m.currChoice().isQuit() {
				return m, tea.Quit
			}
		case key.Matches(msg, selectKeyQuit):
			return m, tea.Quit

		case key.Matches(msg, selectKeyUp):
			m.cursor = (m.cursor - 1 + len(m.choices)) % len(m.choices)

		case key.Matches(msg, selectKeyDown):
			m.cursor = (m.cursor + 1) % len(m.choices)
		}
	}

	return m, nil
}

var (
	selItemRen = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
	headerRen  = lipgloss.NewStyle().Foreground(lipgloss.Color("87")).Align(lipgloss.Center).Padding(0, 2).Bold(true)
	redText    = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true).Align(lipgloss.Right)
)

func (m Model) View() string {
	// The header
	header := headerRen.Render("Welcome to tenp [Terminal NotePad]\nYour Personal Notebook")
	body := ""
	// Iterate over our choices
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = selItemRen.Render(">")
			choice.title = selItemRen.Render(choice.title)
		}
		body += fmt.Sprintf("%s %s\n", cursor, choice.title)
	}

	// The footer
	footer := fmt.Sprintf("Press %s to select. (%s to quit)", selItemRen.Render("[Enter]"), redText.Render("q/ctrl+c"))
	// Send the UI for rendering
	return fmt.Sprintf("%s \n\n%s \n%s\n", header, body, footer)
}
