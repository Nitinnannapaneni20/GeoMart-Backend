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
    CreatedAt     time.Time      `json:"created_at"`
}
