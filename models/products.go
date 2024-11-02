package models

import (
    "time"
)

type Location struct{
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"not null;unique"`
}

type Category struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"not null;unique"`
    LocationID uint   `gorm:"not null"`
}

type ProductType struct {
    ID         uint   `gorm:"primaryKey"`
    Name       string `gorm:"not null;unique"`
    CategoryID uint   `gorm:"not null"`
    LocationID uint   `gorm:"not null"`
}

type ProductData struct {
    ID         uint      `gorm:"primaryKey"`
    Name       string    `gorm:"not null"`
    Brand      string    `gorm:"not null"`
    Quantity   uint      `gorm:"not null"`
    CategoryID uint      `gorm:"not null"`
    TypeID     uint      `gorm:"not null"`
    LocationID  uint      `gorm:"not null"`
    Cost       float64   `gorm:"not null"`
    Description string    `gorm:"not null"`
    Category   Category   `gorm:"foreignKey:CategoryID"`
    ProductType ProductType `gorm:"foreignKey:TypeID"`
    ImageURL   string
    CreatedAt  time.Time `gorm:"autoCreateTime"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
