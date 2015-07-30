package main

import (
	"fmt"

	"code.google.com/p/gopass"
	"github.com/jzelinskie/geddit"
)

type Ui struct {
	TakeTheH int
}

func InitUi() (ui *Ui) {
	ui = &Ui{
		TakeTheH: 1,
	}

	return
}

func (ui *Ui) PrintSubmissions(submissions []*geddit.Submission) {
	for i, sub := range submissions {
		fmt.Printf("%d: Title: %s\nAuthor: %s, Subreddit: %s\nVotes: %d\n\n", i, sub.Title, sub.Author, sub.Subreddit, sub.Score)
		s.Last = sub.FullID
	}
}

func (ui *Ui) CommandlineReadline(prompt string) (output string) {
	fmt.Print(prompt)
	fmt.Scanf("%s", &output)
	return
}

func (ui *Ui) CommandlineSecretInput(prompt string) (output string) {
	output, _ = gopass.GetPass(prompt)

	return
}
