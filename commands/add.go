package commands

import (
    "finance-tracker/models"
    "finance-tracker/storage"
    "fmt"
)

var idCounter = 1

func AddTransaction(date, txType, category, note string, amount float64) {
    tx := models.Transaction{
        ID:       idCounter,
        Date:     date,
        Type:     txType,
        Category: category,
        Amount:   amount,
        Note:     note,
    }
    err := storage.SaveTransaction(tx)
    if err != nil {
        fmt.Println("Erro ao salvar:", err)
        return
    }
    idCounter++
    fmt.Println("Transação adicionada com sucesso.")
}
