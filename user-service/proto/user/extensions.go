package user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate hooks an event lifecycle into GORM so that we
// generate a UUID for our ID column before the entity is saved.
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("Id", uuid.String())
}
