package main

const (
	USER_AGENT = "gotexv0;gedditAgentv1"
)

var s *Session
var ui *Ui

func main() {
	s = InitSession()
	ui = InitUi()

	s.Frontpage(s.Limit, "")
	c := ""
	for c != "q" {
		c = ui.CommandlineReadline(":")
		DoInput(c)
	}
}
