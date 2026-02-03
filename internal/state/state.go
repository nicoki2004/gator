package state

import (
	"github.com/nicoki2004/gator/internal/config"
	"github.com/nicoki2004/gator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}
