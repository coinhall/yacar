package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/coinhall/yacar/internal/yacar"
	"golang.org/x/exp/maps"
)

const IgnoreErrorFile = "ignore_error.txt"

func GetLocalYacarFilePaths(projRoot string) []string {
	fileNames := yacar.GetAllFilesWithExt()

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

func GetErrorFilePath(projRoot string) (errorFilePath string) {
	if err := filepath.Walk(projRoot, func(path string, info os.FileInfo, err error) error {
		if len(errorFilePath) > 0 {
			return nil
		}

		if err != nil {
			log.Printf("error while walking path: %s", err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Base(path) == IgnoreErrorFile {
			errorFilePath = path
		}

		return nil
	}); err != nil {
		panic(fmt.Sprintf("error while walking root dir: %s", err))
	}

	return errorFilePath
}

func GetFileNameNoSuffix(fp, suffix string) string {
	return strings.TrimSuffix(filepath.Base(fp), suffix)
}
