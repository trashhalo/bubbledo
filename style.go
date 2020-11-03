package main

import (
	"github.com/muesli/termenv"
	"github.com/trashhalo/tsx"
)

func style(path []tsx.Path, text string) termenv.Style {
	colors := termenv.ColorProfile()
	node := path[len(path)-1]

	if node.Name == "header" {
		if hasClass(node, "focus") {
			return termenv.String(text).Bold()
		}
	}

	if len(path)-2 >= 0 {
		parent := path[len(path)-2]
		if parent.Name == "task" {
			str := termenv.String(text + " ")
			if hasClass(parent, "focus") {
				str = str.Bold()
			}
			switch node.Name {
			case "priority":
				str = str.Foreground(colors.Color("#9CFFFA"))
			case "project":
				str = str.Foreground(colors.Color("#ACF39D"))
			}
			return str
		}
	}

	return termenv.String(text)
}

func hasClass(node tsx.Path, class string) bool {
	for _, c := range node.ClassNames {
		if class == c {
			return true
		}
	}
	return false
}
