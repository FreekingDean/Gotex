package main

import (
	"github.com/jzelinskie/geddit"
)

type Session struct {
	Type  string
	Last  string
	Limit int

	Session      *geddit.Session
	LoginSession *geddit.LoginSession
}

func InitSession() (s *Session) {
	s = &Session{
		Type:    "std",
		Last:    "",
		Limit:   10,
		Session: geddit.NewSession(USER_AGENT),
	}
	return
}
func (s *Session) Login() {
	username := ui.GetString("username: ")
	password := ui.GetSecret("password: ")

	sesh, err := geddit.NewLoginSession(username, password, USER_AGENT)
	if err != nil {
		return
	}

	s.Type = "login"
	s.LoginSession = sesh
	s.Last = ""
	s.Session = nil
	s.Frontpage(s.Limit, "")
}

func (s *Session) Logout() {
	s.Type = "std"
	if s.LoginSession != nil {
		s.LoginSession.Clear()
		s.LoginSession = nil
	}
	s.Session = geddit.NewSession(USER_AGENT)
	s.Frontpage(s.Limit, "")
}

func (s *Session) Frontpage(limit int, last string) {
	// Set listing options
	subOpts := geddit.ListingOptions{
		Limit: limit,
		After: last,
	}

	var submissions []*geddit.Submission
	if s.Type == "login" {
		submissions, _ = s.LoginSession.Frontpage(geddit.DefaultPopularity, subOpts)
	} else {
		submissions, _ = s.Session.DefaultFrontpage(geddit.DefaultPopularity, subOpts)
	}

	s.Last = ui.PrintSubmissions(submissions)
}
