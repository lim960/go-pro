package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai") //上海

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt *FormatTime    `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *FormatTime    `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type FormatTime time.Time

func (t *FormatTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.In(loc).Format("2006-01-02 15:04:05"))), nil
}

func (t FormatTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.In(loc).Format("2006-01-02 15:04:05"), nil
}

func (t *FormatTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = FormatTime(vt)
	}
	return nil
}

func (u *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("updated_at", time.Now())
	return
}
