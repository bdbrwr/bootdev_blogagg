package main

import "fmt"

func handlerLogin(s *state, cmd command) error { // This will be the function signature of all command handlers!
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Println("User switch succesfully!")
	return nil
}
