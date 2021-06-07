package clotas

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type ClotaFile struct {
	// YYYYMMDD-NNNN-scriptName.sh
	dateStr    string
	number     int				// NNNN - serial of the script within the day
	numberStr  string
	scriptName string			// scriptName - defaults to 'script'
	info       os.FileInfo		// Fileinfo if the file already exists
	Name       string			// filename
	path       string
	extension  string
}

func (ClotaFile) New(scriptName string) *ClotaFile {
	t := time.Now()
	f := new(ClotaFile)

	f.dateStr = time.Now().Format(dateLayout())
	f.number  = 1
	f.numberStr = strconv.Itoa(f.number)
	f.scriptName = scriptName
	f.extension = DefaultFileType

	f.Name = GenerateName(f.scriptName, t, f.number)
	cwd, _ := os.Getwd()
	f.path = filepath.Join(cwd, DefaultTargetFolder, f.Name)

	file, _ := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.Close()
	f.info, _ = os.Stat(f.path)

	return f
}

func (ClotaFile) NewFromFile(file string) *ClotaFile {
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

// GetNext /**
// Generate new ClotaFile object, with incremented serial number from an existing provided file.
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
