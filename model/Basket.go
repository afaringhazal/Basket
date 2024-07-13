package model

import "time"

type Basket struct {
	Basketiden uint32    `gorm:"not null"`
	Username   string    `gorm:"not null"` // Link to the User model
	Data       string    `gorm:"type:varchar(128)"`
	State      bool      `gorm:"type: Boolean"`
	Created_at time.Time `gorm:"type: TIMESTAMP NOT NULL"`
	Update_at  time.Time `gorm:"type: TIMESTAMP"`
}
