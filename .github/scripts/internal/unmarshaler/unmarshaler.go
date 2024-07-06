package unmarshaler

import (
	"encoding/json"
	"os"
)

func UnmarshalInto[T any](path string, container T) (T, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return container, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return container, err
	}

	if err := json.Unmarshal(data, &container); err != nil {
		return container, err
	}

	return container, nil
}
