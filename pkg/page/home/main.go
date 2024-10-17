package home

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Option struct {
	title  string
	keyTag rune
}

type Model struct {
	choices []Option
	cursor  int
}

func InitialModel() Model {
	return Model{
		choices: []Option{
			{title: "✏️  - [A]dd a new note", keyTag: 'a'},
			{title: "📒 - [V]iew all notes", keyTag: 'v'},
			{title: "📝 - [E]dit a note", keyTag: 'e'},
			{title: "🗑️  - [D]elete a note", keyTag: 'd'},
			{title: "🔍 - [S]earch notes by keyword", keyTag: 's'},
			{title: "🚪 - [Q]uit", keyTag: 'q'},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

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
