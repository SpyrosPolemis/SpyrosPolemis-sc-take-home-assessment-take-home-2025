package folder

import (
	"errors"
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Return slice containing all child folders of selected folder

	// Error checking
	if name == "" {
		return []Folder{}, errors.New("name is required")
	}
	if orgID == uuid.Nil {
		return []Folder{}, errors.New("orgID is required")
	}

	// Get child folders
	folders := f.folders
	childFolders := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID && strings.Contains(f.Name, name) {
			childFolders = append(childFolders, f)
		}
	}
	return childFolders, nil
}
