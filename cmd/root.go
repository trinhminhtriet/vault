package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trinhminhtriet/vault/internal"
)

var rootCmd = &cobra.Command{
	Use:   "vault",
	Short: "Vault CLI",
	Long:  `Vault CLI is a command-line tool for managing secrets and sensitive information.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add version flag
	rootCmd.Flags().BoolP("version", "V", false, "Print the version of the application")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if versionFlag, _ := cmd.Flags().GetBool("version"); versionFlag {
			fmt.Printf("Vault version: %s, Date: %s, Commit: %s\n", internal.VERSION, internal.DATE, internal.COMMIT)
			os.Exit(0)
		}
	}
}
