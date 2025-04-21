package models

import (
    "time"
)

type UserData struct {
    ID           uint      `gorm:"primaryKey"`
    Auth0ID      string    `gorm:"uniqueIndex;not null"`
    Name         string
    Email        string    `gorm:"uniqueIndex;not null"`
    Phone        string
    AddressLine1 string
    AddressLine2 string
    City         string
    State        string
    Zip          string
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
