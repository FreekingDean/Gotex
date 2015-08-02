package main

import (
	"flag"

	uipkg "./ui"
)

const (
	USER_AGENT = "gotexv0;gedditAgentv1"
)

var s *Session
var ui uipkg.Ui

func main() {
	var interfaceFlag string
	flag.StringVar(&interfaceFlag, "interface", "ncurse", "Use ncurses over std out [basic,ncurse]")
	flag.Parse()

	s = InitSession()
	ui = uipkg.InitUi(interfaceFlag)

	s.Frontpage(s.Limit, "")
	var c int
	for c != 'q' {
		c = ui.GetCommand()
		DoInput(c)
	}
	ui.KillScreen()
}

func HandleErr(err error) {
	//ui.KillScreen()
	if err != nil {
		panic(err)
	}
}
