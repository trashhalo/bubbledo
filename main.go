package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JamesClonk/go-todotxt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/trashhalo/tsx"
)

func main() {
	tasklist, err := todotxt.LoadFromFilename("todo.txt")
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(model{tasklist: tasklist})
	p.EnterAltScreen()
	defer p.ExitAltScreen()
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type model struct {
	tasklist todotxt.TaskList

	newTodoText string
	editingTodo bool
	focus       int
	err         error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.editingTodo {
			return handleKeyWhileEditing(m, msg)
		}
		return handleKeyWhileNotEditing(m, msg)
	}
	return m, nil
}

func (m model) View() string {
	return app(m).String(style, nil)
}

func app(m model) tsx.T {
	var children []tsx.T
	children = append(children, header(m, m.focus == 0))
	for i, t := range m.tasklist {
		children = append(children, todoItem(t, m.focus != 0 && m.focus-1 == i))
	}

	children = append(children, tsx.T{
		Name: "footer",
		Text: "---\nUp/Down - move around, Space - complete todo, Space - reopen todo",
	})

	return tsx.T{
		Name:     "app",
		Children: children,
	}
}

func header(m model, hasFocus bool) tsx.T {
	text := m.newTodoText
	if !m.editingTodo && text == "" {
		text = "Press Enter to start editing. Enter again to create a new todo. Or Escape to cancel"
	}
	var classNames []string
	if hasFocus {
		classNames = append(classNames, "focus")
	}

	return tsx.T{
		Name:       "header",
		Text:       fmt.Sprintf("[%s]\n", text),
		ClassNames: classNames,
	}
}
