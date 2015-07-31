package main

import (
	"fmt"

	"code.google.com/p/gopass"
	"github.com/jzelinskie/geddit"
)

type Ui interface {
	PrintSubmissions([]*geddit.Submission)
	KillScreen()
	CommandlineReadline(string) string
	CommandlineSecretInput(string) string
}

type BasicUi struct{}

func InitUi(intFaceType string) (ui Ui) {
	if intFaceType == "ncurse" {
		ui = InitNcurseUi()
	} else {
		ui = InitBasicUi()
	}

	return
}

func InitBasicUi() (ui *BasicUi) {
	ui = &BasicUi{}

	return
}

func (ui *BasicUi) PrintSubmissions(submissions []*geddit.Submission) {
	for i, sub := range submissions {
		fmt.Printf("%d: Title: %s\nAuthor: %s, Subreddit: %s\nVotes: %d\n\n", i, sub.Title, sub.Author, sub.Subreddit, sub.Score)
		s.Last = sub.FullID
	}
}

func (ui *BasicUi) KillScreen() {
	fmt.Println("Done.")
}

func (ui *BasicUi) CommandlineReadline(prompt string) (output string) {
	fmt.Print(prompt)
	fmt.Scanf("%s", &output)
	return
}

func (ui *BasicUi) CommandlineSecretInput(prompt string) (output string) {
	output, _ = gopass.GetPass(prompt)

	return
}
