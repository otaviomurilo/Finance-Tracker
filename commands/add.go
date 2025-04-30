package commands

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

func AddTransaction(date, txType, category, note string, amount float64) {
    id := getNextID()

    file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    record := []string{
        strconv.Itoa(id),
        date,
        txType,
        category,
        fmt.Sprintf("%.2f", amount),
        note,
    }

    if err := writer.Write(record); err != nil {
        fmt.Println("Erro ao escrever no arquivo:", err)
    } else {
        fmt.Println("Transação adicionada com sucesso!")
    }
}

func getNextID() int {
    file, err := os.Open("data.csv")
    if err != nil {
        return 1
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return 1
    }

    return len(records) + 1
}
