package commands

import (
	"github.com/mcamack/cli/api"
	"github.com/spf13/cobra"
)

func selectCoin() string {
	promptContent := promptContent{
		"Please choose a coin",
		"What coin would you like the price of?",
	}

	var selectionList = []string{"bitcoin", "ethereum", "cardano"}

	coin := promptGetInput(promptContent, selectionList)
	return coin
}

// priceCmd is a crypto subcommand
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Get the price of a crypto",
	Long:  `Long help message for price subcommand`,
	Run: func(cmd *cobra.Command, args []string) {
		coin := selectCoin()
		api.GetPrice(coin)
	},
}

func init() {
	cryptoCmd.AddCommand(priceCmd)
}
