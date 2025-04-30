package storage

import (
    "encoding/csv"
    "os"
    "strconv"
    "finance-tracker/models"
)

const FilePath = "data.csv"

func SaveTransaction(tx models.Transaction) error {
    file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    record := []string{
        strconv.Itoa(tx.ID),
        tx.Date,
        tx.Type,
        tx.Category,
        strconv.FormatFloat(tx.Amount, 'f', 2, 64),
        tx.Note,
    }
    return writer.Write(record)
}
