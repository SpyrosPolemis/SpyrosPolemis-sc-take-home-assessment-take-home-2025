package folder

import "strings"

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Move folder to destination folder.

	updatedFolders := []Folder{}

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
