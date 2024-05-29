package model

import "time"

type Basket struct {
	Id         int       `gorm:"unique"`
	Username   string    `gorm:"not null"` // Link to the User model
	Data       string    `gorm:"type:varchar(128)" validate:"max=2048"`
	State      bool      `gorm:"type: Boolean"`
	Created_at time.Time `gorm:"type: TIMESTAMP NOT NULL"`
	Update_at  time.Time `gorm:"type: TIMESTAMP"`
}
