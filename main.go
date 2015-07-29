package main

import (
	"fmt"
	"github.com/jzelinskie/geddit"
)

const (
	USER_AGENT = "gotexv0;gedditAgentv1"
)

type Session struct {
	Type  string
	Last  string
	Limit int

	Session      *geddit.Session
	LoginSession *geddit.LoginSession
}

func main() {
	s := InitSession()

	s.frontpage(s.Limit, "")
	var c string
	fmt.Print(":")
	fmt.Scanf("%s", &c)
	for c != "q" {
		if c == "n" {
			s.frontpage(s.Limit, s.Last)
		}
		if c == "login" {
			s.Login()
		}
		if c == "logout" {
			s.Logout()
		}
		fmt.Print(":")
		fmt.Scanf("%s", &c)
	}
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
	var username, password string
	fmt.Print("username: ")
	fmt.Scanf("%s", &username)
	fmt.Print("password: ")
	fmt.Scanf("%s", &password)

	sesh, err := geddit.NewLoginSession(username, password, USER_AGENT)
	if err != nil {
		return
	}

	s.Type = "login"
	s.LoginSession = sesh
	s.Last = ""
}

func (s *Session) Logout() {
	s.Type = "std"
	if s.LoginSession != nil {
		s.LoginSession.Clear()
	}
}

func (s *Session) frontpage(limit int, last string) {
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

	for i, sub := range submissions {
		fmt.Printf("%d: Title: %s\nAuthor: %s, Subreddit: %s\nVotes: %d\n\n", i, sub.Title, sub.Author, sub.Subreddit, sub.Score)
		s.Last = sub.FullID
	}
}
