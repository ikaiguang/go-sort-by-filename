package ossort

import (
	"os"
)

var (
	filenameHandler = NewSortByFilename()
	fileInfoHandler = NewSortByFileInfo()
	sortNameHandler = NewSortByName()
)

// FilenameAsc .
func FilenameAsc(nameSlice []string) {
	filenameHandler.Asc(nameSlice)
}

// FilenameDesc .
func FilenameDesc(nameSlice []string) {
	filenameHandler.Desc(nameSlice)
}

// FileInfoAsc .
func FileInfoAsc(infoSlice []os.FileInfo) {
	fileInfoHandler.Asc(infoSlice)
}

// FileInfoDesc .
func FileInfoDesc(infoSlice []os.FileInfo) {
	fileInfoHandler.Desc(infoSlice)
}

// Asc .
func Asc(infos []Name) {
	sortNameHandler.Asc(infos)
}

// Desc .
func Desc(infos []Name) {
	sortNameHandler.Desc(infos)
}
