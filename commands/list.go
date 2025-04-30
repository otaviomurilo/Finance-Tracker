package commands

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

func ListTransactions() {
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

    fmt.Printf("%-5s %-10s %-8s %-12s %-10s %s\n", "ID", "Data", "Tipo", "Categoria", "Valor", "Nota")
    fmt.Println("---------------------------------------------------------------")

    for _, record := range records {
        id, _ := strconv.Atoi(record[0])
        amount, _ := strconv.ParseFloat(record[4], 64)

        fmt.Printf("%-5d %-10s %-8s %-12s %-10.2f %s\n",
            id, record[1], record[2], record[3], amount, record[5])
    }
}
