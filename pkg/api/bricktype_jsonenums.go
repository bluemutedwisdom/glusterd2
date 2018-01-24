// generated by jsonenums -type=BrickType; DO NOT EDIT

package api

import (
	"encoding/json"
	"fmt"
)

var (
	_BrickTypeNameToValue = map[string]BrickType{
		"Brick":   Brick,
		"Arbiter": Arbiter,
	}

	_BrickTypeValueToName = map[BrickType]string{
		Brick:   "Brick",
		Arbiter: "Arbiter",
	}
)

func init() {
	var v BrickType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BrickTypeNameToValue = map[string]BrickType{
			interface{}(Brick).(fmt.Stringer).String():   Brick,
			interface{}(Arbiter).(fmt.Stringer).String(): Arbiter,
		}
	}
}

// MarshalJSON is generated so BrickType satisfies json.Marshaler.
func (r BrickType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BrickTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BrickType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BrickType satisfies json.Unmarshaler.
func (r *BrickType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BrickType should be a string, got %s", data)
	}
	v, ok := _BrickTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BrickType %q", s)
	}
	*r = v
	return nil
}