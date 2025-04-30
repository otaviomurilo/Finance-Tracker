package cmd

import (
    "finance-tracker/commands"
    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "Lista todas as transações",
    Run: func(cmd *cobra.Command, args []string) {
        commands.ListTransactions()
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
