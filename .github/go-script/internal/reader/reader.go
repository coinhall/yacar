package reader

import (
	"fmt"
	"path/filepath"
)

func AsGlobPattern(projRoot, dirs, files string) string {
	if dirs != "" {
		return fmt.Sprintf(`%s/%s/%s`, projRoot, dirs, files)
	}

	return fmt.Sprintf(`%s/%s`, projRoot, files)
}

func filePathsFromGlob(globPattern string) ([]string, error) {
	return filepath.Glob(globPattern)
}
