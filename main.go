package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/nicoki2004/gator/internal/config"
	"github.com/nicoki2004/gator/internal/database"
	"github.com/nicoki2004/gator/internal/state"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		println(err)
	}

	appState := &state.State{
		Cfg: cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state.State, command) error),
	}

	db, err := sql.Open("postgres", appState.Cfg.DBURL)
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	dbQueries := database.New(db)

	appState.Db = dbQueries

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

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
