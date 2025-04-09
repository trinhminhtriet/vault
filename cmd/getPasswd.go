package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/trinhminhtriet/vault/internal"

	"github.com/spf13/cobra"
)

// getPasswdCmd represents the getPasswd command
var getPasswdCmd = &cobra.Command{
	Use:   "getPasswd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		decryptedPassword := args[0]

		privateKeyPath := os.Getenv("VAULT_PRIVATE_KEY_PATH")
		if privateKeyPath == "" {
			log.Fatalf("Environment variable VAULT_PRIVATE_KEY_PATH is not set")
		}

		password, err := internal.DecryptCode(decryptedPassword, privateKeyPath)
		if err != nil {
			log.Fatalf("Error decrypting password: %v", err)
		}

		fmt.Println(password)
	},
}

func init() {
	rootCmd.AddCommand(getPasswdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPasswdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPasswdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
