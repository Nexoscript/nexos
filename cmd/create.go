package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Erstellt ein neues virtuelles Terminal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		// Kommando zum Erstellen des virtuellen Terminals (PTY)
		fmt.Printf("Erstelle virtuelles Terminal: %s\n", name)

		// Hier kannst du das PTY und den Startbefehl hinzufügen
		err := exec.Command("bash").Start()
		if err != nil {
			fmt.Printf("Fehler beim Erstellen des Terminals: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Virtuelles Terminal '%s' erstellt\n", name)
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
