package main

import (
	"fmt"

	"github.com/JamesClonk/go-todotxt"
	"github.com/trashhalo/tsx"
)

func todoItem(task todotxt.Task, hasFocus bool) tsx.T {
	var classNames []string
	if hasFocus {
		classNames = append(classNames, "focus")
	}

	children := []tsx.T{checkbox(task.Completed)}

	if task.HasPriority() {
		children = append(children, tsx.T{
			Name: "priority",
			Text: fmt.Sprintf("(%s)", task.Priority),
		})
	}

	children = append(children, tsx.T{Name: "text", Text: task.Todo})

	for _, project := range task.Projects {
		children = append(children, tsx.T{
			Name: "project",
			Text: fmt.Sprintf("+%s", project),
		})
	}

	return tsx.T{
		Name:       "task",
		ClassNames: classNames,
		Children:   children,
		Text:       "\n",
	}
}

func checkbox(checked bool) tsx.T {
	txt := " "
	if checked {
		txt = "X"
	}

	return tsx.T{
		Name: "checkbox",
		Text: fmt.Sprintf("[%s]", txt),
	}
}
