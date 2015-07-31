package main

import (
	"flag"
)

const (
	USER_AGENT = "gotexv0;gedditAgentv1"
)

var s *Session
var ui Ui

func main() {
	var interfaceFlag string
	flag.StringVar(&interfaceFlag, "-interface", "ncurse", "Use ncurses over std out")
	flag.Parse()

	s = InitSession()
	ui = InitUi(interfaceFlag)

	s.Frontpage(s.Limit, "")
	c := ""
	for c != "q" {
		c = ui.CommandlineReadline(":")
		DoInput(c)
	}
}

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}
