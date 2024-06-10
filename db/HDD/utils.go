package HDD

import (
	"encoding/json"
	"errors"
)

func unMarshalToField(field []interface{}, to ...interface{}) error {
	if len(field) != len(to) {
		return errors.New("field length is not match to length")
	} else {

		for i, f := range field {
			if err := json.Unmarshal(f.([]byte), to[i]); err != nil {
				return err
			}
		}

		return nil
	}
}
