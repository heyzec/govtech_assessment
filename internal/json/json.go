package json

import (
    "encoding/json"
    "io"
)

func DecodeParams(r io.Reader, view interface{}) error {
    decoder := json.NewDecoder(r)
    err := decoder.Decode(view)
    if err != nil {
        return err
    }
    return nil
}

func EncodeView(view interface{}) ([]byte, error) {
	data, err := json.Marshal(view)
	if err != nil {
		return nil, err
	}
	return data, nil
}

