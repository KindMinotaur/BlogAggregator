package main

import (
	"fmt"
	"log"

	"home/gingersnap/Projects/BlogAggregator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

func handlerLogin(s *state, cmd command) error {
	if cmd.arguments == nil {
		log.Fatalln("Please enter a username")
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		log.Fatalf("Couldn't login user: %v", err)
	}
	fmt.Printf("Logged in user: %v", cmd.arguments)
	return nil
}

type commands struct {
	c map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	// Placeholder comment
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("lane")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
