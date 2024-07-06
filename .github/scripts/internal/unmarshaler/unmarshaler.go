package unmarshaler

import (
	"encoding/json"
	"errors"
	"os"
)

func UnmarshalInto[T any](path string, container T) (T, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return container, nil
	} else if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &container); err != nil {
		return container, err
	}

	return container, nil
}
