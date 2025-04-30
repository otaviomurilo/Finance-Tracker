package cmd

import (
    "finance-tracker/commands"
    "github.com/spf13/cobra"
)

var (
    txType   string
    category string
    amount   float64
    note     string
    date     string
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Adiciona uma nova transação",
    Run: func(cmd *cobra.Command, args []string) {
        commands.AddTransaction(date, txType, category, note, amount)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)

    addCmd.Flags().StringVarP(&txType, "type", "t", "", "Tipo da transação (income ou expense)")
    addCmd.Flags().StringVarP(&category, "category", "c", "", "Categoria da transação")
    addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Valor")
    addCmd.Flags().StringVarP(&note, "note", "n", "", "Nota")
    addCmd.Flags().StringVarP(&date, "date", "d", "", "Data (YYYY-MM-DD)")
}
