package clotas

import (
	"os"
	"regexp"
	"strings"
	"time"
)

func GetFileListForDay(folder string, t time.Time) []string {
	dir, _ := os.Open(folder)
	files, _ := dir.Readdir(0)

	var matchingFiles []string

	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), t.Format(dateLayout())) {
			matchingFiles = append(matchingFiles, file.Name())
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
