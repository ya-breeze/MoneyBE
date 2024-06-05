package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ya-breeze/geekbudgetbe/pkg/generated/goserver"
)

// type AccountNoId struct {
// 	gorm.Model
// 	goserver.AccountNoId

// 	userId string
// }

type Account struct {
	gorm.Model

	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId string    `gorm:"index:idx_user_id"`

	goserver.AccountNoId
}
