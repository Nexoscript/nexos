package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "nexos",
	Short: "Ein Terminal-Multiplexer ähnlich wie screen",
	Long:  `Dies ist ein Beispiel für ein CLI-Programm zum Erstellen und Verwalten von virtuellen Terminals.`,
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	// Hier können globale Flags oder Setup-Code hinzugefügt werden
}
