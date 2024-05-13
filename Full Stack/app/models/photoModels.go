package models

type Photo struct {
    ID       uint   `gorm:"primary_key" json:"id"`
    Title    string `json:"title"`
    Caption  string `json:"caption"`
    PhotoURL string `json:"photo_url"`
    UserID   uint   `json:"user_id"`
}
