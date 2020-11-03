package main

import (
	"github.com/JamesClonk/go-todotxt"
	tea "github.com/charmbracelet/bubbletea"
)

func handleKeyWhileEditing(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.editingTodo = false
		m.newTodoText = ""
		return m, nil
	case "enter":
		if m.focus == 0 {
			return handleEnterOnHeaderEditing(m)
		}
		return handleEnterOnTaskEditing(m)
	default:
		m.newTodoText += string(msg.Rune)
		return m, nil
	}
}

func handleKeyWhileNotEditing(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "down":
		if m.focus+1 <= len(m.tasklist) {
			m.focus++
		}
		return m, nil
	case "up":
		if m.focus != 0 {
			m.focus--
		}
		return m, nil
	case "enter":
		if m.focus == 0 {
			return handleEnterOnHeaderNotEditing(m)
		}
		return handleEnterOnTaskNotEditing(m)
	case " ":
		if m.focus > 0 {
			task := m.tasklist[m.focus-1]
			if task.Completed {
				task.Reopen()
			} else {
				task.Complete()
			}
			m.tasklist[m.focus-1] = task
		}
		return m, nil
	default:
		return m, nil
	}
}

func handleEnterOnHeaderNotEditing(m model) (tea.Model, tea.Cmd) {
	m.newTodoText = ""
	m.editingTodo = true
	return m, nil
}

func handleEnterOnHeaderEditing(m model) (tea.Model, tea.Cmd) {
	m.editingTodo = false

	task, err := todotxt.ParseTask(m.newTodoText)
	if err != nil {
		m.err = err
	} else {
		m.tasklist.AddTask(task)
	}

	m.newTodoText = ""
	return m, nil
}

func handleEnterOnTaskNotEditing(m model) (tea.Model, tea.Cmd) {
	return m, nil
}

func handleEnterOnTaskEditing(m model) (tea.Model, tea.Cmd) {
	return m, nil
}
