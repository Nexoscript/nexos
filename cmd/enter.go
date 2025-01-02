package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var enterCmd = &cobra.Command{
	Use:   "enter [name]",
	Short: "Betritt ein virtuelles Terminal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		// Hier kannst du den Terminalzugriff für das entsprechende Terminal implementieren
		fmt.Printf("Betrete virtuelles Terminal: %s\n", name)

		// PTY starten (z. B. Bash oder ein anderes Kommando)
		err := exec.Command("bash").Start()
		if err != nil {
			fmt.Printf("Fehler beim Betreten des Terminals: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Im virtuellen Terminal '%s' eingeloggt\n", name)
	},
}

func init() {
	RootCmd.AddCommand(enterCmd)
}
