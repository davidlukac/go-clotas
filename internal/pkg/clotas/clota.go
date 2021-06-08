package clotas

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ClotaFile struct {
	// YYYYMMDD-NNNN-scriptName.sh
	dateStr    string
	number     int // NNN - serial of the script within the day
	numberStr  string
	scriptName string      // scriptName - defaults to 'script'
	info       os.FileInfo // Fileinfo if the file already exists
	Name       string      // filename
	path       string
	extension  string
}

// New
// Create new ClotaFile object from basic provided info. No actual file is created.
// @param scriptName string	(Optional) If empty, defaults to DefaultScriptName.
// @param n int 			(Optional) If negative number is provided defaults to 1.
// @param t *time.Time		(Optional) If empty, set to Now; used for the date in the file name.
func (ClotaFile) New(scriptName string, n int, t *time.Time) *ClotaFile {
	var tt time.Time

	if t == nil {
		tt = time.Now()
	} else {
		tt = *t
	}

	f := new(ClotaFile)

	f.dateStr = tt.Format(dateLayout())

	if n < 0 {
		f.number = 1
	} else {
		f.number = n
	}

	f.numberStr = fmt.Sprintf("%03d", f.number)

	if scriptName == "" {
		scriptName = DefaultScriptName
	}
	f.scriptName = scriptName

	f.extension = DefaultFileType

	f.Name = GenerateName(f.scriptName, tt, f.number)

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

// GetNextFromList /**
// Generate new ClotaFile object from provided list of ClotaFile-s with incremented serial.
func (ClotaFile) GetNextFromList(files []ClotaFile, scriptName string) *ClotaFile {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	return ClotaFile{}.GetNext(files[len(files)-1], scriptName)
}

// GetNext /**
// Generate new ClotaFile object, with incremented serial number from an existing provided file.
func (ClotaFile) GetNext(file ClotaFile, scriptName string) *ClotaFile {
	return ClotaFile{}.New(scriptName, file.number+1, nil)
}

// CreateFile /**
// Create actual file on the file system and update the object with information.
func (cf *ClotaFile) CreateFile() {
	cwd, _ := os.Getwd()
	cf.path = filepath.Join(cwd, DefaultTargetFolder, cf.Name)

	file, _ := os.OpenFile(cf.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, DefaultFileMode)
	file.Close()
	cf.info, _ = os.Stat(cf.path)
}
