package clotas

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func GetFileListForDay(folder string, t time.Time) []ClotaFile {
	dir, _ := os.Open(folder)
	files, _ := dir.Readdir(0)

	var matchingFiles []ClotaFile

	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), t.Format(dateLayout())) {
			matchingFiles = append(matchingFiles, *ClotaFile{}.NewFromFileInfo(file))
		}
	}

	dir.Close()

	return matchingFiles
}

func GetFiles(folder string) []string {
	dir, _ := os.Open(folder)
	files, _ := dir.Readdir(0)

	var matchingFiles []string

	re := regexp.MustCompile(`^\d+.*`)

	for _, file := range files {
		if !file.IsDir() &&
			strings.HasSuffix(file.Name(), "."+DefaultFileType) &&
			re.MatchString(file.Name()) {
			matchingFiles = append(matchingFiles, file.Name())
		}
	}

	dir.Close()

	return matchingFiles
}

// AssureClotasDir
// Make sure Clotas folder exists.
func AssureClotasDir(folder string) {
	cwd, _ := os.Getwd()

	if folder == "" {
		folder = filepath.Join(cwd, DefaultTargetFolder)
	}

	_, err := os.Stat(folder)

	if err != nil {
		os.Mkdir(folder, 0755)
	}
}
