package authority

import (
	"database/sql/driver"
	"errors"
)

type UserIDColumn struct {
	ID interface{}
}

func (id *UserIDColumn) GormDataType() string {
	return idColumnType
}

// Scan - loads value from database.
func (id *UserIDColumn) Scan(src interface{}) error {
	switch src := src.(type) {
	case uint:
	case string:
		id.ID = src
	case nil:
		return errors.New("unsupported value for id column in database")
	}

	return nil
}

// Value - set value in database.
func (id UserIDColumn) Value() (driver.Value, error) {
	switch id.ID.(type) {
	case uint:
	case string:
		return id.ID, nil
	}
	return nil, errors.New("unsupported value for id column")
}
