package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/trinhminhtriet/vault/internal"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
)

// mfaCmd represents the mfa command
var mfaCmd = &cobra.Command{
	Use:   "mfa",
	Short: "Decrypt and generate OTP from TOTP",
	Long:  `Decrypt and generate OTP from TOTP`,
	Run: func(cmd *cobra.Command, args []string) {
		decryptedTotpValue := args[0]

		privateKeyPath := os.Getenv("VAULT_PRIVATE_KEY_PATH")
		if privateKeyPath == "" {
			log.Fatalf("Environment variable VAULT_PRIVATE_KEY_PATH is not set")
		}

		totpValue, err := internal.DecryptCode(decryptedTotpValue, privateKeyPath)
		if err != nil {
			log.Fatalf("Error decrypting TOTP: %v", err)
		}

		otp, err := totp.GenerateCode(totpValue, time.Now())
		if err != nil {
			log.Fatalf("Error generating OTP: %v", err)
		}

		fmt.Println(otp)
	},
}

func init() {
	rootCmd.AddCommand(mfaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mfaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mfaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
