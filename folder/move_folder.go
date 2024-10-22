package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Move folder to destination folder.

	var folderToMove *Folder
	var destinationFolder *Folder

	// Check for existence of source and destination folders
	for _, folder := range f.folders {
		if folder.Name == name {
			folderToMove = &folder
		}
		if folder.Name == dst {
			destinationFolder = &folder
		}
	}

	// Error checking
	if folderToMove == nil {
		return nil, errors.New("source folder does not exist")
	}
	if destinationFolder == nil {
		return nil, errors.New("destination folder does not exist")
	}
	if folderToMove.Name == dst {
		return nil, errors.New("cannot move a folder to itself")
	}
	if strings.HasPrefix(folderToMove.Paths, dst+".") {
		return nil, errors.New("cannot move a folder to a child of itself")
	}
	if folderToMove.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}

	updatedFolders := []Folder{}

	// Move folder
	for _, f := range f.folders {
		if f.Name == name {
			f.Paths = dst + "." + name
		} else if strings.HasPrefix(f.Paths, name+".") {
			f.Paths = dst + "." + f.Paths
		}
		updatedFolders = append(updatedFolders, f)
	}
	return updatedFolders, nil
}
