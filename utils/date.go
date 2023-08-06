package utils

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Date time.Time

var _ json.Unmarshaler = &Date{}

const dateFormat = "2006-01-02"
const dateTimeFormat = "2006-01-02T15:04"

func (d *Date) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*d = Date(time.Time{})
		return nil
	}

	var formatTries = []string{dateFormat, dateTimeFormat, time.RFC3339}
	var t time.Time
	for _, f := range formatTries {
		t, err = time.ParseInLocation(f, s, time.UTC)
		if err == nil {
			break
		}
	}
	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	if t.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return json.Marshal(time.Time(*d))
}

func (d *Date) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		*d = Date(t)
	}
	return nil
}

func (d *Date) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}
	return time.Time(*d), nil
}
