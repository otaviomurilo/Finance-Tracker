package cmd

import (
    "finance-tracker/commands"
    "github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
    Use:   "balance",
    Short: "Mostra o saldo atual",
    Run: func(cmd *cobra.Command, args []string) {
        commands.ShowBalance()
    },
}

func init() {
    rootCmd.AddCommand(balanceCmd)
}
