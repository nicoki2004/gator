// Package state contains the logic for managing the application's global state,
// including user configuration and the database connection.
package state

import (
	"github.com/nicoki2004/gator/internal/config"
	"github.com/nicoki2004/gator/internal/database"
)

// State represents the synchronized state of the CLI.
// It holds the configuration loaded from the JSON file and the database client.
type State struct {
	Db       *database.Queries
	Cfg      *config.Config
	Commands any
}
