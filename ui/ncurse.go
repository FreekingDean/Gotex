package ui

import (
	"github.com/rthornton128/goncurses"

	"github.com/jzelinskie/geddit"
)

type NcurseUi struct {
	contents *goncurses.Window
	titleBar *goncurses.Window
	menu     *goncurses.Menu
}

func (ui *NcurseUi) initColors() {
	if goncurses.HasColors() {
		err := goncurses.StartColor()
		HandleErr(err)

		goncurses.UseDefaultColors()
		goncurses.InitPair(1, goncurses.C_WHITE, -1)
		goncurses.InitPair(2, goncurses.C_BLACK, goncurses.C_BLUE)

		goncurses.InitPair(3, goncurses.C_YELLOW, -1)
	}
}

func InitNcurseUi() (ui *NcurseUi) {
	ui = &NcurseUi{}

	var err error
	ui.contents, err = goncurses.Init()
	HandleErr(err)

	inputY, inputX := ui.contents.MaxYX()

	ui.initColors()

	ui.contents.Resize(inputY-1, inputX)
	ui.contents.MoveWindow(1, 0)
	ui.contents.Keypad(true)

	ui.titleBar, err = goncurses.NewWindow(1, inputX, 0, 0)
	HandleErr(err)
	ui.titleBar.ColorOn(2)
	ui.titleBar.SetBackground(goncurses.Char('.') | goncurses.ColorPair(2))

	return
}

func (ui *NcurseUi) KillScreen() {
	goncurses.End()
}

func freeMenu(menu *goncurses.Menu) {
	if menu != nil {
		menu.UnPost()
		for _, item := range menu.Items() {
			item.Free()
		}
		menu.Free()
	}
}

func (ui *NcurseUi) PrintSubmissions(submissions []*geddit.Submission) (last string) {
	freeMenu(ui.menu)
	ui.titleBar.Printf("Frontpage")
	ui.titleBar.Refresh()

	ui.contents.Clear()
	menu_items := make([]*goncurses.MenuItem, len(submissions))
	for i, sub := range submissions {
		//fmt.Sprintf("%d: Title: %s\n   Author: %s, Subreddit: %s\n   Votes: %d\n\n", i, sub.Title, sub.Author, sub.Subreddit, sub.Score)
		item, err := goncurses.NewItem(sub.Title, sub.FullID)
		err = nil
		HandleErr(err)

		menu_items[i] = item
		last = sub.FullID
	}

	if len(menu_items) > 0 {
		var err error
		ui.menu, err = goncurses.NewMenu(menu_items)
		ui.menu.SetSpacing(2, 3, 1)
		ui.menu.SetBackground(goncurses.ColorPair(3))
		HandleErr(err)

		ui.menu.Post()
	}

	ui.contents.Refresh()

	return
}

func (ui *NcurseUi) GetCommand() (output int) {
	goncurses.Update()
	goncurses.Echo(false)
	output = int(ui.contents.GetChar())
	goncurses.Echo(true)

	return
}

func (ui *NcurseUi) Println(msg string) {
	ui.contents.Println(msg)
	ui.contents.Refresh()
}

func (ui *NcurseUi) GetString(prompt string) (output string) {
	output = string(ui.contents.GetChar())

	return
}

func (ui *NcurseUi) GetSecret(prompt string) (output string) {
	goncurses.Echo(false)
	output = ui.GetString(prompt)
	goncurses.Echo(true)

	return
}

func (ui *NcurseUi) Handle(input int) {
	if action, ok := goncurses.DriverActions[goncurses.Key(input)]; ok && ui.menu != nil {
		ui.menu.Driver(action)
	}
}
