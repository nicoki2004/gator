package main

import "fmt"

func requireArgs(cmd command, count int) error {
	if len(cmd.Args) != count {
		return fmt.Errorf("usage: %s requires %d argument(s), got %d", cmd.Name, count, len(cmd.Args))
	}
	return nil
}
