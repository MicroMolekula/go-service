package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

type Message struct {
	Id      uint      `json:"id" gorm:"primary_key"`
	UserId  uint      `json:"user_id" gorm:"not null"`
	Date    time.Time `json:"date"`
	Context string    `json:"context"`
	Type    bool      `json:"isGPT"`
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), "\"")
	dateTime, err := ParseDate(timeString)
	if err != nil {
		return err
	}
	*t = DateTime(dateTime)
	return nil
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*t).Format("02-01-2006 15:04:05"))), nil
}

func (t *DateTime) ToTime() time.Time {
	return time.Time(*t)
}

func ParseDate(date string) (time.Time, error) {
	timeLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.Now(), err
	}
	return time.ParseInLocation("02-01-2006 15:04", date, timeLocation)
}

func DateToString(t time.Time) string {
	return t.Format("02-01-2006 15:04")
}

func (t *DateTime) Value() (driver.Value, error) {
	return t.ToTime(), nil // Конвертируем в стандартный time.Time
}

func (t *DateTime) Scan(src interface{}) error {
	if src == nil {
		*t = DateTime(time.Time{})
		return nil
	}
	if tim, ok := src.(time.Time); ok {
		*t = DateTime(tim)
		return nil
	}
	return fmt.Errorf("cannot convert %v to DateTime", src)
}
