package models

type Transaction struct {
    ID       int
    Date     string  // Ex: "2025-04-30"
    Type     string  // "income" ou "expense"
    Category string
    Amount   float64
    Note     string
}
