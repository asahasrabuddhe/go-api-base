package models

import (
	"time"
)

type Model struct {
	Id        int        `json:"id" db:"name:id"`
	CreatedOn *time.Time `json:"created_on" db:"name:created_on"`
	UpdatedOn *time.Time `json:"updated_on" db:"name:updated_on"`
	DeletedOn *time.Time `json:"deleted_on,omitempty" db:"name:deleted_on"`
}
