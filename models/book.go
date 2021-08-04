package models

import ( // import dependencies
	"time"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MediumPrice float32   `json:"mediumPrice"`
	ImageURL    string    `json:"img_url"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
	DeletedAt   time.Time `json:"deleted"`
}
