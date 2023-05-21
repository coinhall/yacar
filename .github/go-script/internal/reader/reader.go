package reader

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/coinhall/yacar/internal/enums"
)

func GetLocalYacarFiles(projRoot string) []string {
	chains := enums.GetAllChainNames()
	files := enums.GetAllFileNames()

	fpMap := make(map[string]bool)

	if err := filepath.Walk(projRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error while walking path: %s", err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		for _, chain := range chains {
			for _, file := range files {
				if strings.Contains(path, chain) && strings.HasSuffix(path, file+".json") {
					fpMap[path] = true
				}
			}
		}

		return nil
	}); err != nil {
		log.Fatalf("error while walking root dir: %s", err)
	}

	filePaths := make([]string, 0, len(fpMap))
	for fp := range fpMap {
		filePaths = append(filePaths, fp)
	}

	return filePaths
}
