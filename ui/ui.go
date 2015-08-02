package ui

import (
	"github.com/jzelinskie/geddit"
)

type Ui interface {
	KillScreen()

	Println(string)
	PrintSubmissions([]*geddit.Submission) string

	GetCommand() int
	GetString(string) string
	GetSecret(string) string

	Handle(int)
	//HandleErr(error)
}

func InitUi(intFaceType string) (ui Ui) {
	if intFaceType == "ncurse" {
		ui = InitNcurseUi()
	} else {
		ui = InitBasicUi()
	}

	return
}

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}
