package main

func DoInput(input int) {
	switch input {
	case 'n':
		s.Frontpage(s.Limit, s.Last)
	case 'l':
		s.Login()
	case 'o':
		s.Logout()
	default:
		ui.Handle(input)
	}
}
