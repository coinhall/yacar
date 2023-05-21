package reader

import (
	"fmt"
)

func AsGlobPattern(projRoot, dirs, files string) string {
	if dirs != "" {
		return fmt.Sprintf(`%s/{%s}/{%s}.json`, projRoot, dirs, files)
	}

	return fmt.Sprintf(`{%s}/{%s}.json`, projRoot, files)
}
