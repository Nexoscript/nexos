package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close [name]",
	Short: "Schließt ein virtuelles Terminal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		// Hier kannst du die Logik implementieren, um das Terminal zu beenden
		fmt.Printf("Schließe virtuelles Terminal: %s\n", name)

		// Beispiel: Beenden eines virtuellen Terminals (z. B. Bash-Prozess)
		err := exec.Command("pkill", "bash").Run()
		if err != nil {
			fmt.Printf("Fehler beim Schließen des Terminals: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Virtuelles Terminal '%s' geschlossen\n", name)
	},
}

func init() {
	RootCmd.AddCommand(closeCmd)
}
