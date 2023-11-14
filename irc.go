package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/lrstanley/girc"
)

func run() {
	client := girc.New(girc.Config{
		Server: "irc.eu.libera.chat",
		Port:   6697,
		Nick:   "NRCBot",
		User:   "NRCBot",
		Name:   "NRCBot",
		SSL:    true,
		Debug:  os.Stdout,
	})

	client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
		c.Cmd.Join("#nrc")
	})

	client.Handlers.Add(girc.PRIVMSG, func(c *girc.Client, e girc.Event) {
		if e.Source.Name == "freethink" && strings.HasPrefix(e.Last(), "!doi") {
			search := strings.Split(strings.TrimSpace(e.Last()), "!doi")
			log.Println("SRCH", search, len(search), len(search[1]))
			if len(search[1]) <= 5 {
				log.Println("!DOI!Input! ")
				c.Cmd.ReplyTo(e, "uh-oh! (Search term needs to be 5 chars or more)")
				return
			}

			doi, err := GetDOI(search[1])
			if err != nil {
				log.Println("!DOI!CrossRef! ", err)
				c.Cmd.ReplyTo(e, "uh-oh!")
				return
			}

			res, err := NewGHIssue(GHIssue{
				Title:  doi.Message.Title[0],
				Body:   doi.Message.Abstract,
				Owner:  "PMaynard",
				Repo:   "NRCBot",
				Labels: []string{"Ideas"},
			})
			if err != nil {
				log.Println("!DOI!GitHub! ", err)
				c.Cmd.ReplyTo(e, "uh-oh!")
				return
			}
			c.Cmd.ReplyTo(e, "Added: "+doi.Message.Title[0]+" ("+res.HTMLURL+")")
			return
		}
	})

	for {
		if err := client.Connect(); err != nil {
			log.Printf("error: %s", err)

			log.Println("reconnecting in 30 seconds...")
			time.Sleep(30 * time.Second)
		} else {
			return
		}
	}
}
