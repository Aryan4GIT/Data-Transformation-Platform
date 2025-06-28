package utils

import (
	"encoding/json"
	"fmt"
	"io"
)


func StreamTransformJSON(r io.Reader, w io.Writer, transform func(key string, value interface{}) (string, interface{})) error {
	dec := json.NewDecoder(r)

	t, err := dec.Token()
	if err != nil || t != json.Delim('{') {
		return fmt.Errorf("expected start of object: %v", err)
	}
	w.Write([]byte("{"))
	first := true
	for dec.More() {
		keyToken, err := dec.Token()
		if err != nil {
			return err
		}
		key := keyToken.(string)
		var value interface{}
		if err := dec.Decode(&value); err != nil {
			return err
		}
		newKey, newValue := transform(key, value)
		if !first {
			w.Write([]byte(","))
		}
		first = false
		keyBytes, _ := json.Marshal(newKey)
		valueBytes, _ := json.Marshal(newValue)
		w.Write(keyBytes)
		w.Write([]byte(":"))
		w.Write(valueBytes)
	}
	t, err = dec.Token()
	if err != nil || t != json.Delim('}') {
		return fmt.Errorf("expected end of object: %v", err)
	}
	w.Write([]byte("}"))
	return nil
}
