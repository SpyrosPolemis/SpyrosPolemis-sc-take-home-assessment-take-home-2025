package folder

import (
	"github.com/gofrs/uuid"
	"strings"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Return slice containing all child folders of selected folder
	folders := f.folders
	childFolders := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID && strings.Contains(f.Name, name) {
			childFolders = append(childFolders, f)
		}
	}
	return childFolders
}
