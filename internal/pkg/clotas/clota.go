package clotas

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ClotaFile struct {
	Name       string
	path       string
	info       os.FileInfo
	extension  string
	dateStr    string
	numberStr  string
	number     int
	scriptName string
}

func New(file string) *ClotaFile {
	f := new(ClotaFile)

	if filepath.IsAbs(file) {
		f.path = file
		_, f.Name = filepath.Split(file)
	} else {
		cwd, _ := os.Getwd()
		f.path = filepath.Join(cwd, DefaultTargetFolder, file)
		f.Name = file
	}

	return f
}

func (ClotaFile) NewFromFileInfo(file os.FileInfo) *ClotaFile {
	f := new(ClotaFile)

	cwd, _ := os.Getwd()

	f.info = file
	f.Name = file.Name()
	f.path = filepath.Join(cwd, DefaultTargetFolder, f.Name)

	parts := strings.Split(f.Name, ".")
	fileName := parts[0]
	f.extension = parts[1]

	// @TODO: In future we might need to be able to handle multiple separators.
	parts = strings.Split(fileName, DefaultSeparator)
	if len(parts) >= 3 {
		f.dateStr = parts[0]
		f.numberStr = parts[1]
		f.scriptName = parts[2]

		number, err := strconv.Atoi(f.numberStr)

		if err != nil {
			panic(fmt.Sprintf("Invalid file name '%s' - unable to parse script number '%s' "+
				"as integer! Expected format is YYYYMMDD-nnn-name.sh.", f.Name, f.numberStr))
		}

		f.number = number
	} else {
		panic(fmt.Sprintf("Invalid file name '%s'! Expected format is YYYYMMDD-nnn-name.sh.", f.Name))
	}

	return f
}

func (ClotaFile) GetNextFromList(files []ClotaFile, scriptName string) *ClotaFile {
	return ClotaFile{}.GetNext(files[len(files)-1], scriptName)
}

func (ClotaFile) GetNext(file ClotaFile, scriptName string) *ClotaFile {
	next := new(ClotaFile)
	next.extension = DefaultFileType
	next.number = file.number + 1
	next.numberStr = fmt.Sprintf("%03d", next.number)
	next.dateStr = file.dateStr

	if len(scriptName) == 0 {
		scriptName = DefaultScriptName
	}

	next.scriptName = scriptName
	next.Name = fmt.Sprintf("%s%s%s%s%s.%s",
		next.dateStr, DefaultSeparator, next.numberStr, DefaultSeparator, next.scriptName, DefaultFileType)

	return next
}
