package ui

import (
	"github.com/rthornton128/goncurses"

	"fmt"
	"github.com/jzelinskie/geddit"
)

const (
	MENU_WIDTH = 5
)

type NcurseUi struct {
	contents *goncurses.Window
	titleBar *goncurses.Window
	menuBar  *goncurses.Window

	menu *goncurses.Menu
}

type menuItem struct {
	id string
}

var menuMap map[int]*menuItem

func (ui *NcurseUi) initColors() {
	if goncurses.HasColors() {
		err := goncurses.StartColor()
		HandleErr(err)

		goncurses.UseDefaultColors()

		//Main window
		goncurses.InitPair(1, goncurses.C_WHITE, -1)
		goncurses.InitPair(2, goncurses.C_RED, -1)
		goncurses.InitPair(3, goncurses.C_GREEN, -1)
		goncurses.InitPair(4, goncurses.C_CYAN, -1)

		//Title colors
		goncurses.InitPair(10, goncurses.C_BLACK, goncurses.C_BLUE)

		//Menu colors
		goncurses.InitPair(20, goncurses.C_YELLOW, -1)
	}
}

func InitNcurseUi() (ui *NcurseUi) {
	ui = &NcurseUi{}

	var err error
	ui.contents, err = goncurses.Init()
	HandleErr(err)

	y, x := ui.contents.MaxYX()

	ui.initColors()

	ui.contents.Resize(y-1, x-MENU_WIDTH)
	ui.contents.MoveWindow(1, MENU_WIDTH)
	ui.contents.Keypad(true)

	ui.titleBar, err = goncurses.NewWindow(1, x, 0, 0)
	HandleErr(err)
	ui.menuBar, err = goncurses.NewWindow(y-1, MENU_WIDTH, 1, 0)
	HandleErr(err)

	ui.menuBar.Border(' ', '|', ' ', ' ', ' ', '|', ' ', ' ')

	ui.titleBar.ColorOn(10)
	ui.titleBar.SetBackground(goncurses.Char('.') | goncurses.ColorPair(10))

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

func (ui *NcurseUi) Println(msg string) {
	ui.contents.Println(msg)
	ui.contents.Refresh()
}

func (ui *NcurseUi) printPos(msg string, Y, X int) {
	curY, curX := ui.contents.CursorYX()
	ui.contents.Move(Y, X)
	ui.contents.Print(msg)
	ui.contents.Move(curY, curX)
	ui.contents.Refresh()
}

func (ui *NcurseUi) PrintSubmissions(submissions []*geddit.Submission) (last string) {
	freeMenu(ui.menu)
	ui.titleBar.Printf("Frontpage")
	ui.titleBar.Refresh()

	ui.contents.Clear()
	menuItems := make([]*goncurses.MenuItem, len(submissions))
	menuMap = make(map[int]*menuItem)
	for i, sub := range submissions {
		//fmt.Sprintf("%d: Title: %s\n   Author: %s, Subreddit: %s\n   Votes: %d\n\n", i, sub.Title, sub.Author, sub.Subreddit, sub.Score)
		menuMap[i] = &menuItem{
			id: sub.FullID,
		}
		item, err := goncurses.NewItem(fmt.Sprintf("%d", i), "")
		HandleErr(err)
		ui.contents.ColorOn(1)
		ui.Println(sub.Title)
		if sub.Score > 0 {
			ui.contents.ColorOn(3)
		} else {
			ui.contents.ColorOn(2)
		}
		ui.Println(fmt.Sprintf("[%5d]", sub.Score))
		ui.contents.ColorOn(4)
		ui.Println(sub.Subreddit + " - " + sub.Author)
		ui.contents.ColorOn(1)

		menuItems[i] = item
		last = sub.FullID
	}

	if len(menuItems) > 0 {
		var err error
		ui.menu, err = goncurses.NewMenu(menuItems)
		HandleErr(err)

		ui.menu.SetSpacing(0, 3, 0)
		ui.menu.SetBackground(goncurses.ColorPair(20))
		ui.menu.SetWindow(ui.menuBar)

		ui.menu.Post()
	}

	ui.menuBar.Refresh()
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
	ui.menuBar.Refresh()
	ui.contents.Refresh()
}
