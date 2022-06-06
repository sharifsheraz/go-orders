package util

import (
	"time"
)

func MillisInLocToTime(millis uint, locName string) (*time.Time, error) {
	loc, err := time.LoadLocation(locName)
	if err != nil {
		return nil, err
	}
	_, locOffset := time.Now().In(loc).Zone()
	_, myOffset := time.Now().Zone()

	t := time.UnixMilli(int64(millis)).Add((time.Duration(-locOffset) + time.Duration(myOffset)) * time.Second)

	return &t, nil
}
