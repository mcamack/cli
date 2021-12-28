package commands

import (
	"github.com/spf13/cobra"
)

func selectCryptoSubcommand() string {
	promptContent := promptContent{
		"Choose a crypto subcommand",
		"What subcommand would you like?",
	}

	var selectionList = []string{"price"}

	chosenSubcommand := promptGetInput(promptContent, selectionList)
	return chosenSubcommand
}

// cryptoCmd represents the crypto command
var cmdFlag = "crypto"
var cryptoCmd = &cobra.Command{
	Use:   cmdFlag,
	Short: "Interactive crypto commands",
	Long:  `Interactive crypto commands`,
	Run: func(cmd *cobra.Command, args []string) {
		subcmd := selectCryptoSubcommand()

		switch subcmd {
		case "price":
			rootCmd.SetArgs([]string{cmdFlag, subcmd})
			rootCmd.Execute()
		}

		rootCmd.SetArgs([]string{cmdFlag})
		rootCmd.Execute()
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(cryptoCmd)
}
