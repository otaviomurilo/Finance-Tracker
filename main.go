package main

import (
    "fmt"
    "finance-tracker/commands"
)

func main() {
    fmt.Println("Bem-vindo ao seu gerenciador de finanças!")
    fmt.Println("Digite a opção: (1) Adicionar, (2) Listar, (3) Saldo")

    var option int
    fmt.Scan(&option)

    switch option {
    case 1:
        var date, txType, category, note string
        var amount float64
        fmt.Println("Data (YYYY-MM-DD):")
        fmt.Scan(&date)
        fmt.Println("Tipo (income/expense):")
        fmt.Scan(&txType)
        fmt.Println("Categoria:")
        fmt.Scan(&category)
        fmt.Println("Valor:")
        fmt.Scan(&amount)
        fmt.Println("Nota:")
        fmt.Scan(&note)

        commands.AddTransaction(date, txType, category, note, amount)
    case 2:
        fmt.Println("Listagem ainda não implementada.")
    case 3:
        fmt.Println("Cálculo de saldo ainda não implementado.")
    default:
        fmt.Println("Opção inválida.")
    }
}
