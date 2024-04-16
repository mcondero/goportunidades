package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint `json:""`
	CreatedAt time.Time
	UpdatedAt time.Time
	Role
	Company
	Location
	Remote
	Link
	Salary
}
