package main

import (
	lp "github.com/charmbracelet/lipgloss"
)

var (
	MenuBorder = lp.NewStyle().Border(lp.RoundedBorder(), true, true).
			BorderForeground(lp.Color("228")).
			PaddingTop(1).
			Align(lp.Center).
			Width(50)

	Title  = lp.NewStyle().Bold(true).Foreground(lp.Color("#F8F8F8")).Align(lp.Center)
	Option = lp.NewStyle().Foreground(lp.Color("86")).Align(lp.Left)

	InputFieldLabels = lp.NewStyle().Border(lp.DoubleBorder(),false, true).
				BorderLeft(true).
				BorderRight(false).
				BorderForeground(lp.Color("#2a9d8f")).
				Foreground(lp.Color("#a3b18a"))
)
