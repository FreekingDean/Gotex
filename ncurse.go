package main

import (
	"fmt"
	"github.com/rthornton128/goncurses"

	"code.google.com/p/gopass"
	"github.com/jzelinskie/geddit"
)

type NcurseUi struct {
	Contents    *goncurses.Window
	CommandLine *goncurses.Window
}

func InitNcurseUi() (ui *NcurseUi) {
	ui = &NcurseUi{}

	var err error
	ui.Contents, err = goncurses.Init()
	HandleErr(err)

	inputY, inputX := ui.Contents.MaxYX()
	ui.CommandLine, err = goncurses.NewWindow(1, inputX, inputY-1, 0)
	HandleErr(err)

	ui.Contents.Resize(inputY-1, inputX)
	ui.Contents.ScrollOk(true)
	ui.CommandLine.Println(":")
	ui.CommandLine.Refresh()

	return
}

func (ui *NcurseUi) KillScreen() {
	goncurses.End()
}

func (ui *NcurseUi) PrintSubmissions(submissions []*geddit.Submission) {
	for i, sub := range submissions {
		fmt.Printf("%d: Title: %s\nAuthor: %s, Subreddit: %s\nVotes: %d\n\n", i, sub.Title, sub.Author, sub.Subreddit, sub.Score)
		s.Last = sub.FullID
	}
}

func (ui *NcurseUi) CommandlineReadline(prompt string) (output string) {
	fmt.Print(prompt)
	fmt.Scanf("%s", &output)
	return
}

func (ui *NcurseUi) CommandlineSecretInput(prompt string) (output string) {
	output, _ = gopass.GetPass(prompt)

	return
}
