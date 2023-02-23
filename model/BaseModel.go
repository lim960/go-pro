package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"pro/util"
	"time"
)

type BaseModel struct {
	ID        uint        `gorm:"primary_key"`
	CreatedAt *FormatTime `gorm:"autoCreateTime"`
	UpdatedAt *FormatTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
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

func (u *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("updated_at", util.FormatSec(time.Now()))
	return
}
