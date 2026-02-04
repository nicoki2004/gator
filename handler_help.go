package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/nicoki2004/gator/internal/state"
)

func handlerHelp(s *state.State, cmd command) error {
	fmt.Println("\n🐊 GATOR RSS READER - HELP")
	fmt.Println("Usage: gator <command> [arguments]")
	fmt.Println()

	// 1. Intentamos recuperar los comandos del State (que es tipo any)
	// Hacemos el "type assertion" al mapa de definiciones
	cmds, ok := s.Commands.(map[string]commandDefinition)
	if !ok {
		return fmt.Errorf("error: could not retrieve command list from state")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintln(w, "COMMAND\tDESCRIPTION\tUSAGE")
	fmt.Fprintln(w, "-------\t-----------\t-----")

	// 2. Obtener y ordenar llaves
	keys := make([]string, 0, len(cmds))
	for k := range cmds {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 3. Imprimir
	for _, name := range keys {
		def := cmds[name]
		fmt.Fprintf(w, "%s\t%s\t%s\n", def.name, def.description, def.usage)
	}

	w.Flush()
	return nil
}
