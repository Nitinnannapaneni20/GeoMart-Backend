package models

import (
    "time"
)

type UserData struct {
    ID        uint           `gorm:"primaryKey"`
    Auth0ID   string         `gorm:"uniqueIndex;not null"`
    Name      string
    Email     string         `gorm:"uniqueIndex;not null"`
    Address   string
    Phone     string
    CreatedAt time.Time
    UpdatedAt time.Time
}
