package main

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(s *state, cmd command)
}

func (c commands) run(s *state, cmd command) error {
	return nil
}
