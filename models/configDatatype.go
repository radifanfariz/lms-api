package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JSONB []map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

// Duration lets us convert between a bigint in Postgres and time.Duration
// in Go
type Duration time.Duration

// Value converts the PgDuration into a string.
func (d Duration) Value() (driver.Value, error) {
	return time.Duration(d * Duration(time.Second)).String(), nil
}

// Scan converts the received string in the format hh:mm:ss into a Duration.
func (d *Duration) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		// Convert format of hh:mm:ss into format parseable by time.ParseDuration()
		v = strings.Replace(v, ":", "h", 1)
		v = strings.Replace(v, ":", "m", 1)
		v += "s"
		dur, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		*d = Duration(dur.Seconds())
		return nil
	default:
		return fmt.Errorf("cannot sql.Scan() PgDuration from: %#v", v)
	}
}
