package models

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := ct.Format(TimeFormat)
	return []byte(`"` + formatted + `"`), nil
}

func (ct *CustomTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 2 {
		return nil
	}
	data = data[1 : len(data)-1]
	ct.Time, err = time.Parse(TimeFormat, string(data))
	return
}

func (ct CustomTime) Value() (driver.Value, error) {
	if ct.IsZero() {
		return nil, nil
	}
	return ct.Time, nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		ct.Time = v
		return nil
	case string:
		var err error
		ct.Time, err = time.ParseInLocation(TimeFormat, v, time.Local)
		if err != nil {
			ct.Time, err = time.Parse(time.RFC3339, v)
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return nil
	}
}
