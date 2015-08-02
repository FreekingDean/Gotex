package ui

import (
	"fmt"

	"code.google.com/p/gopass"
	"github.com/jzelinskie/geddit"
)

type BasicUi struct{}

func InitBasicUi() (ui *BasicUi) {
	ui = &BasicUi{}

	return
}

func (ui *BasicUi) PrintSubmissions(submissions []*geddit.Submission) (last string) {
	for i, sub := range submissions {
		fmt.Printf("%d: Title: %s\nAuthor: %s, Subreddit: %s\nVotes: %d\n\n", i, sub.Title, sub.Author, sub.Subreddit, sub.Score)
		last = sub.FullID
	}
	return
}

func (ui *BasicUi) KillScreen() {
	fmt.Println("Done.")
}

func (ui *BasicUi) GetCommand() (output int) {
	//output = ' '
	output_s := ui.GetString(":")
	if len(output_s) > 0 {
		output = int(output_s[0])
	}

	return
}

func (ui *BasicUi) Println(msg string) {
	fmt.Println(msg)
}

func (ui *BasicUi) GetString(prompt string) (output string) {
	fmt.Print(prompt)
	fmt.Scanf("%s", &output)
	return
}

func (ui *BasicUi) GetSecret(prompt string) (output string) {
	output, _ = gopass.GetPass(prompt)

	return
}

func (ui *BasicUi) Handle(input int) {
	fmt.Print(input)
}
