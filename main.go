package main

import (
	"fmt"
	"log"
	"os"

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
	if f, ok := c.c[cmd.name]; ok {
		return f(s, cmd)
	}
	log.Fatalf("Unknown command: %v", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.c == nil {
		c.c = make(map[string]func(*state, command) error)
	} else if _, exists := c.c[name]; exists {
		log.Fatalf("Command %v already registered", name)
	}
	c.c[name] = f
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Couldn't read config: %v", err)
	}

	s := &state{
		config: &cfg,
	}
	
	cmds := &commands{}
	cmds.register("login", handlerLogin)
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalln("Please provide a command")
	}
	cmd := command{
		name:      args[0],
		arguments: args[1:],
	}
	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatalf("Command %v failed: %v", cmd.name, err)
	}
}
