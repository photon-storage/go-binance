package common

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type StringInt64 string

func (s *StringInt64) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) > 0 && b[0] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		*s = StringInt64(str)
		return nil
	}

	var num int64
	if err := json.Unmarshal(b, &num); err != nil {
		return err
	}
	*s = StringInt64(strconv.FormatInt(num, 10))
	return nil
}
