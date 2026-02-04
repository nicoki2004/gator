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
		log.Fatal("Failed to read config: check ~/.gatorconfig.json exists and is valid")
	}

	appState := &state.State{
		Cfg: cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]commandDefinition),
	}

	db, err := sql.Open("postgres", appState.Cfg.DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	dbQueries := database.New(db)

	appState.Db = dbQueries
	appState.Commands = cmds.registeredCommands

	// --- User Management ---
	cmds.register(commandDefinition{
		name:        "register",
		description: "Create a new user and set it as active",
		usage:       "register <username>",
		minArgs:     1,
		handler:     handlerRegister,
	})

	cmds.register(commandDefinition{
		name:        "login",
		description: "Set the current active user in the config file",
		usage:       "login <username>",
		minArgs:     1,
		handler:     handlerLogin,
	})

	cmds.register(commandDefinition{
		name:        "users",
		description: "List all registered users in the database",
		usage:       "users",
		minArgs:     0,
		handler:     handlerUsers,
	})

	// --- Feed Management ---
	cmds.register(commandDefinition{
		name:        "addfeed",
		description: "Add a new RSS feed to the system",
		usage:       "addfeed <name> <url>",
		minArgs:     2,
		handler:     middlewareLoggedIn(handlerAddFeed),
	})

	cmds.register(commandDefinition{
		name:        "feeds",
		description: "Show all available feeds and their owners",
		usage:       "feeds",
		minArgs:     0,
		handler:     handlerFeeds,
	})

	// --- Follow Management ---
	cmds.register(commandDefinition{
		name:        "follow",
		description: "Follow an existing feed by its URL",
		usage:       "follow <url>",
		minArgs:     1,
		handler:     middlewareLoggedIn(handlerFollow),
	})

	cmds.register(commandDefinition{
		name:        "following",
		description: "List all feeds you are currently following",
		usage:       "following",
		minArgs:     0,
		handler:     middlewareLoggedIn(handlerFollowing),
	})

	cmds.register(commandDefinition{
		name:        "unfollow",
		description: "Stop following a feed by its URL",
		usage:       "unfollow <url>",
		minArgs:     1,
		handler:     middlewareLoggedIn(handlerUnfollow),
	})

	// --- System & Aggregation ---
	cmds.register(commandDefinition{
		name:        "agg",
		description: "Periodically fetch and store posts from all feeds",
		usage:       "agg <time_between_reqs>",
		minArgs:     1,
		handler:     handlerAgg,
	})

	cmds.register(commandDefinition{
		name:        "browse",
		description: "View recent posts from feeds you follow",
		usage:       "browse [limit]",
		minArgs:     0,
		handler:     middlewareLoggedIn(handlerBrowse),
	})

	cmds.register(commandDefinition{
		name:        "reset",
		description: "Clear all users from the database (DANGER!)",
		usage:       "reset",
		minArgs:     0,
		handler:     handlerReset,
	})

	cmds.register(commandDefinition{
		name:        "help",
		description: "Show this help menu",
		usage:       "help",
		minArgs:     0,
		handler:     handlerHelp,
	})

	// cmds.register("register", handlerRegister)
	// cmds.register("reset", handlerReset)
	// cmds.register("users", handlerUsers)
	// cmds.register("agg", handlerAgg)
	// cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	// cmds.register("feeds", handlerFeeds)
	// cmds.register("follow", middlewareLoggedIn(handlerFollow))
	// cmds.register("following", middlewareLoggedIn(handlerFollowing))
	// cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	// cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	//
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
