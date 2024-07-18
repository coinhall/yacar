package writer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile[T any](path string, data []T) error {
	var sb strings.Builder
	sbEnc := json.NewEncoder(&sb)
	sbEnc.SetEscapeHTML(false)
	sbEnc.SetIndent("", "  ")
	if err := sbEnc.Encode(data); err != nil {
		panic(err)
	}

	parentDir := filepath.Dir(path)
	if err := os.MkdirAll(parentDir, 0o755); err != nil {
		return err
	}

	// If file exists, overwrite the contents completely
	if err := os.WriteFile(path, []byte(sb.String()), 0o644); err != nil {
		return err
	}

	return nil
}
