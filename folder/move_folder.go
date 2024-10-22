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
	if strings.Contains(destinationFolder.Paths, folderToMove.Name) {
		return nil, errors.New("cannot move a folder to a child of itself")
	}
	if folderToMove.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}

	updatedFolders := []Folder{}

	destinationPath := destinationFolder.Paths
	if destinationPath != "" {
		destinationPath += "."
	}
	sourcePath := folderToMove.Paths

	// Move folder
	for _, f := range f.folders {
		if f.Name == name {
			f.Paths = destinationPath + name
		} else if strings.HasPrefix(f.Paths, sourcePath) {
			f.Paths = destinationPath + name + "." + f.Name
		}
		updatedFolders = append(updatedFolders, f)
	}
	return updatedFolders, nil
}
