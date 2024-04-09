package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    		uuid.UUID `json:"id" gorm:"primaryKey"`
	CreatedAt 	time.Time `json:"created_at"`
	Name  		string    `json:"name"`
	Email 		string    `json:"email"`
}