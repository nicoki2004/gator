package main

import (
	"log"
	"os"

	"github.com/nicoki2004/gator/internal/config"
	"github.com/nicoki2004/gator/internal/state"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		println(err)
	}

	appState := &state.State{
		Cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state.State, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	currentCmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = cmds.run(appState, currentCmd)
	if err != nil {
		log.Fatal(err)
	}
}
