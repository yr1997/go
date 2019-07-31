package models

type Post struct {
    Id  uint64 `gorm:"primary_key" json:"id"`
    Body string `sql:"size:255" json:"body"`
}