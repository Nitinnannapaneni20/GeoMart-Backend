package models

import (
    "time"
    "gorm.io/datatypes"
)

type Order struct {
    ID            uint           `gorm:"primaryKey" json:"id"`
    UserID        string         `gorm:"index" json:"user_id"`
    Items         datatypes.JSON `gorm:"type:jsonb;not null" json:"items"`
    TotalAmount   float64        `gorm:"not null" json:"total_amount"`
    Currency      string         `gorm:"default:GBP" json:"currency"`
    PaymentStatus string         `gorm:"default:COMPLETED" json:"payment_status"`
    TransactionID string         `gorm:"unique;not null" json:"transaction_id"`
    Name          string         `gorm:"size:255" json:"name"`
    Email         string         `gorm:"size:255" json:"email"`
    Phone         string         `gorm:"size:50" json:"phone"`
    AddressLine1  string         `gorm:"size:255" json:"addressLine1"`
    AddressLine2  string         `gorm:"size:255" json:"addressLine2"`
    City          string         `gorm:"size:100" json:"city"`
    State         string         `gorm:"size:100" json:"state"`
    Zip           string         `gorm:"size:20" json:"zip"`
    CreatedAt     time.Time      `json:"created_at"`
}
