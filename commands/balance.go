package commands

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

func ShowBalance() {
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Erro ao ler o CSV:", err)
        return
    }

    var income, expense float64

    for _, record := range records {
        value, _ := strconv.ParseFloat(record[4], 64)
        if record[2] == "income" {
            income += value
        } else if record[2] == "expense" {
            expense += value
        }
    }

    balance := income - expense

    fmt.Printf("Receitas: R$ %.2f\n", income)
    fmt.Printf("Despesas: R$ %.2f\n", expense)
    fmt.Printf("Saldo atual: R$ %.2f\n", balance)
}
