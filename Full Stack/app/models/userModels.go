package models

import "time"

type User struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-"`
    Photos    []Photo   `json:"photos"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
