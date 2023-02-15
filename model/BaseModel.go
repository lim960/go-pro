package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt *FormatTime
	UpdatedAt *FormatTime
	DeletedAt *FormatTime `sql:"index"`
}

type FormatTime time.Time

func (t *FormatTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t FormatTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *FormatTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = FormatTime(vt)
	}
	return nil
}
