package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

func GetLocalYacarFilePaths(projRoot string) []string {
	fileNames := GetAllFilesWithExt()

	fpSet := make(map[string]struct{})

	if err := filepath.Walk(projRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error while walking path: %s", err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		for _, fileName := range fileNames {
			if strings.HasSuffix(path, fileName) {
				fpSet[path] = struct{}{}
			}
		}

		return nil
	}); err != nil {
		panic(fmt.Sprintf("error while walking root dir: %s", err))
	}

	filePaths := maps.Keys(fpSet)
	slices.Sort(filePaths)

	return filePaths
}
